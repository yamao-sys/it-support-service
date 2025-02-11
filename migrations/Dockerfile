# syntax=docker/dockerfile:1

FROM golang:1.23.1-alpine3.20

WORKDIR /app

COPY . .

RUN go mod tidy && \
	go install github.com/rubenv/sql-migrate/...@latest && \
	go install github.com/joho/godotenv/cmd/godotenv@latest

RUN apk add --update make
