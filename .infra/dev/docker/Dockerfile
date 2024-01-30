FROM golang:1.21.0 

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on  

RUN go install github.com/pressly/goose/v3/cmd/goose@latest