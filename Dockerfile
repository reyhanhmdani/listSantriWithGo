FROM golang:1.20-alpine AS builder

LABEL authors="rey"

WORKDIR /app

COPY go.mod go.sum ./

ENTRYPOINT ["top", "-b"]