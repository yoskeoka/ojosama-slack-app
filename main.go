package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jiro4989/ojosama"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func main() {
	webApi := slack.New(
		os.Getenv("SLACK_BOT_TOKEN"),
		slack.OptionAppLevelToken(os.Getenv("SLACK_APP_TOKEN")),
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
	)
	socketClient := socketmode.New(
		webApi,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "sm: ", log.Lshortfile|log.LstdFlags)),
	)
	authTest, authTestErr := webApi.AuthTest()
	if authTestErr != nil {
		fmt.Fprintf(os.Stderr, "SLACK_BOT_TOKEN is invalid: %v\n", authTestErr)
		os.Exit(1)
	}
	selfUserId := authTest.UserID
	socketClient.Debugf("SelfUserID: %v", selfUserId)

	subscriber := SlackEventSubscriber{
		BotUser:      authTest,
		WebApi:       webApi,
		SocketClient: socketClient,
	}

	go func() {
		for event := range socketClient.Events {
			err := subscriber.HandleSocketModeEvent(event)
			if err != nil {
				log.Printf("Handle Event error: %v", err)
			}
		}
	}()

	socketClient.Run()
}

type SlackEventSubscriber struct {
	BotUser      *slack.AuthTestResponse
	SocketClient *socketmode.Client
	WebApi       *slack.Client
}

func (es *SlackEventSubscriber) HandleSocketModeEvent(envelop socketmode.Event) error {
	switch envelop.Type {
	case socketmode.EventTypeConnecting:
		fmt.Println("Connecting to Slack with Socket Mode...")
	case socketmode.EventTypeConnectionError:
		fmt.Println("Connection failed. Retrying later...")
	case socketmode.EventTypeConnected:
		fmt.Println("Connected to Slack with Socket Mode.")
	case socketmode.EventTypeDisconnect:
		fmt.Println("Disconnected from Slack with Socket Mode.")
	case socketmode.EventTypeSlashCommand:
		es.SocketClient.Ack(*envelop.Request,
			map[string]interface{}{
				// slash command typed by the user will be visible.
				"response_type": "in_channel",
			},
		)

		cmd, ok := envelop.Data.(slack.SlashCommand)
		if !ok {
			return fmt.Errorf("ignored %+v", envelop.Data)
		}

		convertedMsg, err := ojosama.Convert(cmd.Text, nil)
		if err != nil {
			return fmt.Errorf("ojosama convert error: %v", err)
		}

		userProfile, err := es.SocketClient.GetUserProfile(&slack.GetUserProfileParameters{
			UserID: cmd.UserID,
		})
		if err != nil {
			return fmt.Errorf("failed to reply: %v", err)
		}

		_, _, err = es.SocketClient.PostMessageContext(context.TODO(),
			cmd.ChannelID,
			slack.MsgOptionCompose(
				slack.MsgOptionText(
					convertedMsg,
					false,
				),
				slack.MsgOptionUsername(fmt.Sprintf("%s (%s)", es.BotUser.User, userProfile.DisplayName)),
			))
		if err != nil {
			return fmt.Errorf("failed to reply: %v", err)
		}

	default:
		es.SocketClient.Debugf("skipped: %+v", envelop)
	}

	return nil
}
