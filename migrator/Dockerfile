FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go install github.com/pressly/goose/v3/cmd/goose@latest