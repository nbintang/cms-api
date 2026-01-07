FROM golang:1.22-alpine AS builder

# Install git (dibutuhkan go mod)
RUN apk add --no-cache git

WORKDIR /app
