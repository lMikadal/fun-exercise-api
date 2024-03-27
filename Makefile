all:
	@go run main.go

test:
	@go test -v ./...

docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down

.PHONY: all test docker-up docker-down