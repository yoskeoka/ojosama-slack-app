export CGO_ENABLED=0

.PHONY: start
start:
	export $$(grep -v '^#' .env | xargs);\
	go run *.go
