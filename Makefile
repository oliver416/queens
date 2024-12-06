#!make
.PHONY: run test swagger format

run:
	go run queens

test:
	go test queens/tests


swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	~/go/bin/swag init -o ./app/docs

format:
	go fmt ./...