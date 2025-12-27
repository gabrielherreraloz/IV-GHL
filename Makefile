install:
	go mod tidy

check:
	go fmt ./...

test:
	go test ./...