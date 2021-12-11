SHELL   := /bin/bash
APP_PORT := ${or ${APP_PORT}, "8080"}

build:
	@go build -v ./...

run:
	@APP_PORT=${APP_PORT} go run github.com/cosmtrek/air@latest

.PHONY: lint
lint:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --fix

go-gen:
	@go generate ./...

# docker

.PHONY: up
up:
	@docker-compose up -d --build

.PHONY: stop
stop:
	@docker-compose stop

.PHONY: down
down:
	@docker-compose down

.PHONY: logs
logs:
	@docker-compose logs -f
