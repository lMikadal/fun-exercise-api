FROM golang:1.22.1-alpine as build-base

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test -v ./...

RUN go build -o ./out/go-fun-exercise .

# ====================

FROM alpine:3.16.2

WORKDIR /app

COPY .env .

COPY --from=build-base /app/out/go-fun-exercise /app/go-fun-exercise

CMD ["/app/go-fun-exercise"]