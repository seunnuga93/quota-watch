BINARY=bin/quota-watch

.PHONY: build run test tidy
build:
	go build -o $(BINARY) .

run:
	go run . --threshold 0.8

test:
	go test ./...

tidy:
	go mod tidy
