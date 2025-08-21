.PHONY: help install run build tests

help:  ##Shows all available commands
	@awk 'BEGIN {FS = ":.*##"; printf "\n\033[32m Makefile\033[0m\n\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

install: ##Wakes up all the needed docker containers
	@docker-compose up -d --build

up: ##Wakes up the application without rebuilding it
	@docker compose up -d

run: ##Executes the app locally
	@go run ./src/user-interface/cmd/main.go

build: ##Compiles the app locally
	@go build ./src/user-interface/cmd/main.go

tests: ##Execute all the tests
	@go test ./tests/...

dclp:
	@sh ./scripts/dclp.sh

shell: ##Opens an interactive shell
	@docker compose exec -it app sh

stop: ##Stop the dockers
	@docker compose down

down: ## Stop the dockers to shut down the system
	@docker compose down -v --rmi all --remove-orphans