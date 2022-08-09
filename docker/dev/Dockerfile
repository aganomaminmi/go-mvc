FROM golang:1.18

ENV GOPATH /go
WORKDIR /

COPY ./go.mod ./go.mod ./
RUN go mod download

# dbmate
RUN curl -fsSL -o /usr/local/bin/dbmate https://github.com/amacneil/dbmate/releases/latest/download/dbmate-linux-amd64
RUN chmod +x /usr/local/bin/dbmate
