FROM golang:1.26.0-alpine3.23 AS builder

WORKDIR /app

COPY . .

WORKDIR /app/services/stack

RUN CGO_ENABLED=0 GOOS=linux go build -o /stack ./cmd/

FROM alpine:3.23
COPY --from=builder /stack /bin/stack

CMD ["/bin/stack"]
