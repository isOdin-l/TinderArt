FROM golang:1.26.0-alpine3.23 AS builder

WORKDIR /app

COPY . .

WORKDIR /app/services/profile

RUN CGO_ENABLED=0 GOOS=linux go build -o /profile ./cmd/

FROM alpine:3.23
COPY --from=builder /profile /bin/profile

CMD ["/bin/profile"]
