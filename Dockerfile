FROM golang:1.10.2-alpine3.7
MAINTAINER Ed King <ed@teddyking.co.uk>

WORKDIR /go/src/github.com/teddyking/snowflake
COPY . .

RUN go build cmd/server/snowflake.go

ENTRYPOINT ./snowflake

ENV PORT 2929
EXPOSE 2929
