.PHONY: all build

all: run

build:
	go build -o ./.bin/auth main.go

run: build
	./.bin/auth

