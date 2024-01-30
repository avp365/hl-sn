#---Build stage---
FROM golang:1.21.0 

WORKDIR /app

ADD go.mod /go/src/
ADD go.sum /go/src/
RUN go mod download

COPY . /go/src/


RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-w -s' -o /go/bin/app


CMD /go/bin/app
