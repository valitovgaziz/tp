FROM golang:1.22.5

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o bin/api cmd/main.go

ENTRYPOINT [ "bin/api" ]