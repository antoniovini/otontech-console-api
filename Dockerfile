FROM golang:alpine

#MAINTAINER Maintainer

ENV GIN_MODE=release
ENV PORT=8000

WORKDIR /go/src/go-docker-dev.to

COPY . .

RUN apk update && apk add --no-cache git
RUN go get .

RUN go build -o ./app ./main.go

EXPOSE $PORT

ENTRYPOINT ["./app"]