DOCKER_TAG = go-fun-exercise

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

docker-build:
	@docker build -t $(DOCKER_TAG):latest .

docker-run:
	docker run $(DOCKER_TAG)

.PHONY: all run test docker-up docker-down swag docker-build docker-run