FROM golang:1.18 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
CMD [ "/app/main" ]
