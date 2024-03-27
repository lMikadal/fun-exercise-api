all:
	@go run main.go

test:
	@go test -v ./...

docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down

swag:
	@swag init

.PHONY: all run test docker-up docker-down swag