FROM golang:1.14.2

RUN apt-get update && \
    apt-get upgrade -y

WORKDIR /go/src/til-review-app

COPY ./go.mod .
COPY ./go.sum .

RUN go get -d -v ./...