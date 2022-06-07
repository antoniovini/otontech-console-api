FROM golang:alpine

#MAINTAINER Maintainer

ENV GIN_MODE=release
ENV PORT=8000

WORKDIR /go/src/go-docker-dev.to

COPY . /go/src/go-docker-dev.to

RUN apk update && apk add --no-cache git
RUN go get .

RUN go build go-docker-dev.to/app

EXPOSE $PORT

ENTRYPOINT ["./app"]