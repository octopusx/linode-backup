FROM golang:1.15-alpine

RUN apk add git make musl-dev gcc

WORKDIR /go/src/github.com/octopusx/linode-backup
COPY . /go/src/github.com/octopusx/linode-backup

RUN go get