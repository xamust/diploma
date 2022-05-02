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


.DEFAULT_GOAL := run
