FROM golang:1.22.5

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o bin/api.exe cmd/main.go

RUN 

ENTRYPOINT [ "bin/api.exe" ]