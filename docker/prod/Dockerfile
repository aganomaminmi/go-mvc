FROM golang:1.18 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GO111MODULE=on go build -o server cmd/server.go cmd/router.go

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/server .
COPY .env .
CMD [ "/app/server" ]
