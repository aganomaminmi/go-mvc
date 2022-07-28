FROM golang:1.18 as go
ENV GOPATH /go
WORKDIR /
COPY ./go.mod ./go.mod ./
RUN go mod download
CMD ["go", "run", "cmd/server.go", "cmd/router.go"]

FROM amacneil/dbmate as migration
RUN echo "now migration setting"
ENV DATABASE_URL mysql://user:password@mysql:3306/go_mvc_db
WORKDIR /
# RUN dbmate migrate
CMD ["dbmate", "migrate"]
