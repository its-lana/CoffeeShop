# step 1: build executable binary
FROM golang:1.20-alpine AS builder
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /coffee-shop-service

# step 2: build a small image
FROM alpine:latest
WORKDIR /app
COPY --from=builder coffee-shop-service .
COPY .env .
RUN mkdir /app/logs
EXPOSE 8083
CMD ["./coffee-shop-service"]