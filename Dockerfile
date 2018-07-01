FROM golang:1.10.3-alpine3.7
MAINTAINER Ed King <ed@teddyking.co.uk>

WORKDIR /go/src/github.com/teddyking/snowflake
COPY . .

RUN go build cmd/snowflake/snowflake.go
RUN go build cmd/snowflakeweb/snowflakeweb.go
RUN mv snowflake* /usr/local/bin/

