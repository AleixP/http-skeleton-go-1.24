.PHONY: help install run test dclp

help:  ##Shows all available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

install: ##Wakes up all the needed docker containers
	docker-compose up --build

run: ##Executes the app locally
	go run ./src/user-interface/cmd/main.go

build: ##Compiles the app locally
	go build ./src/user-interface/cmd/main.go

tests: ##Execute all the tests
	go test ./tests/...

dclp:
	sh ./scripts/dclp.sh