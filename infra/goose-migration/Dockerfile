FROM golang:1.26-alpine AS builder

RUN apk add --no-cache git ca-certificates

RUN go install github.com/pressly/goose/v3/cmd/goose@latest


FROM alpine:3.23

WORKDIR /migrations
COPY --from=builder /go/bin/goose /usr/local/bin/goose
