.PHONY: init build run test

all:
		make init
		make build
		make run
init:
		go mod tidy
build:
		make init
#		go build -v ./cmd/app
		go build -o build/server -v ./cmd/app
run:
		make init
		go run ./cmd/app
pretest:
		go install honnef.co/go/tools/cmd/staticcheck@latest
		sudo apt install golint
test:
		make pretest
		golint ./...
		go vet ./...
		staticcheck ./...

.DEFAULT_GOAL := run
