# ojosama-slack-app

Slack App ですわ

## 準備

```sh
cp .env.example .env
```

Slack Appを作成して、Slack App Tokenを発行しまして、`SLACK_APP_TOKEN=`の値を置き換えるて下さいまし。
同じく、Workspace installをしまして、発行したSlack Bot Tokenで`SLACK_BOT_TOKEN=`の値を置き換えるて下さいまし。

## Heroku デプロイ

```sh
make deploy-heroku
```
