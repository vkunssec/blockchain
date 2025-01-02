MAKEFLAGS += -j2
.DEFAULT_GOAL := run

blockchain:
	go run cmd/main.go


env:
	cp .env.example .env

build:
	make swagger
	go build -ldflags="-s -w" -o tmp/main.exe cmd/main.go

target: run-blockchain

run:
	make build
	make target

run-blockchain:
	./tmp/blockchain

swagger:
	swag init -g cmd/main.go --parseDependency --parseInternal

dev:
	go get -u github.com/vkunssec/husky
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	go get -u github.com/swaggo/swag/cmd/swag
	make swagger
	air server

all: swagger dev
