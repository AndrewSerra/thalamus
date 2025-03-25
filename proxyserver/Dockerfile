FROM golang:1.24.1-bookworm as builder

WORKDIR /app

COPY go.mod .
COPY . .

RUN go mod download

RUN go build -o proxyserver main.go
