FROM golang:1.18

ENV GOPATH /go
WORKDIR /

COPY go.* ./
RUN go mod download \
  && go install github.com/cosmtrek/air@v1.27.3

