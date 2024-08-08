FROM golang:1.22.5

RUN go get github.com/pressly/goose/cmd/goose


WORKDIR /app

COPY migrations/* app/migrations/

CMD [ "goose", "up", "-dir=/app/migrations" ]