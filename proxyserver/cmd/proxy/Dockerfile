FROM golang:1.24.1-bookworm AS builder

WORKDIR /app

COPY ../../go.mod .
COPY . .

RUN go mod download

RUN go build -o proxyserver cmd/proxy/main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/proxyserver /app/proxyserver

CMD ["/app/proxyserver"]
