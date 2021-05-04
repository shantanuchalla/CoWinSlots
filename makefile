.PHONY: build
build:
	go build -o slot-checker cmd/main/main.go

.PHONY: run
run:
	./slot-checker

.PHONY: execute
execute: build run
