#!make
.PHONY: run test

run:
	go run queens

test:
	go test queens/tests

