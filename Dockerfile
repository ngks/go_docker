FROM golang:1.16-alpine

RUN mkdir /app
WORKDIR /app
RUN go mod init test/module
RUN go get github.com/gin-gonic/gin
