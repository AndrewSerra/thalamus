FROM golang:1.24.1-bookworm AS builder

WORKDIR /app

COPY go.mod .
COPY . .

RUN go mod download

RUN go build -o analyticsserver main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/analyticsserver /app/analyticsserver

CMD ["/app/analyticsserver"]
