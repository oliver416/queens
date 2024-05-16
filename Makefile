#!make
.PHONY: run test

run:
	go run main.go

test:
	go test queens/tests

