.DEFAULT_GOAL := build
BIN_FILE=transtop

install:
	go mod tidy

test:
	go test ./...

check:
	go fmt ./...

build:
	@go build -o "${BIN_FILE}"