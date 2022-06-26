export CGO_ENABLED=0

.PHONY: start
start:
	export $$(grep -v '^#' .env | xargs);\
	go run *.go

.PHONY: build
build:
	go build -o bin/ojosama-slack-app \
		-tags 'netgo' -installsuffix netgo

.PHONY: deploy-heroku
deploy-heroku:
	if heroku whoami 2>&1 | grep "not logged in"; then heroku login; fi
	heroku container:login
	export $$(grep -v '^#' .env | xargs);\
    heroku container:push -a $${HEROKU_APP_NAME} worker;\
	heroku config:set -a $${HEROKU_APP_NAME} SLACK_APP_TOKEN=$${SLACK_APP_TOKEN} SLACK_BOT_TOKEN=$${SLACK_BOT_TOKEN};\
	heroku container:release -a $${HEROKU_APP_NAME} worker
