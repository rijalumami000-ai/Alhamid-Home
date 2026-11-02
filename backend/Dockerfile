# Multi-stage Dockerfile for Go Backend
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server ./cmd/server

FROM alpine:latest

WORKDIR /app
RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Jakarta

COPY --from=builder /app/server .
COPY .env.example .env

RUN mkdir -p /app/uploads

EXPOSE 8088

CMD ["./server"]

