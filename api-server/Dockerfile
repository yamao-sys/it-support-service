# syntax=docker/dockerfile:1

FROM golang:1.24.0-alpine3.20

WORKDIR /app

COPY . .

RUN go mod tidy && \
	go install honnef.co/go/tools/cmd/staticcheck@latest && \
	go install github.com/joho/godotenv/cmd/godotenv@latest && \
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest && \
	go install github.com/air-verse/air@latest

RUN apk add --update make
