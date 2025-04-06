FROM golang:1.24.1-bookworm AS builder

WORKDIR /app

COPY ../../go.mod .
COPY cmd/registration/. .

RUN go mod download

RUN go build -o registrationserver .

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/registrationserver /app/registrationserver

CMD ["/app/registrationserver"]
