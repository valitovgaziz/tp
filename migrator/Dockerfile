FROM golang:1.22.5

WORKDIR /app

COPY migrations/* app/migrations/

COPY go.mod app/

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

CMD ["goose", "-dir=/app/migrations", "user=${DB_USER}", "dbname=${DB_NAME}", "sslmode=${SSLmode}" "up"]