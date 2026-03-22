FROM golang:1.26.0-alpine3.23 AS builder

WORKDIR /app

COPY . .

WORKDIR /app/services/swipe

RUN CGO_ENABLED=0 GOOS=linux go build -o /swipe ./cmd/

FROM alpine:3.23
COPY --from=builder /swipe /bin/swipe

CMD ["/bin/swipe"]
