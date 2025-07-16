# Stage 1: Build
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

# Установим git (требуется для некоторых go get)
RUN apk add --no-cache git

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o bot .

# Stage 2: Final image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bot .

# Установим корневой сертификат (если телеграм требует TLS)


# Запуск
CMD ["./bot"]
