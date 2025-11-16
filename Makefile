.DEFAULT_GOAL := build
BIN_FILE=transtop

install:
	go mod tidy

test:
	go test ./...

check:
	go fmt ./...
	go vet ./...

clean:
	go clean
	rm --force "cp.out"
	rm --force nohup.out

build:
	@go build -o "${BIN_FILE}"
	
run:
	./"${BIN_FILE}"