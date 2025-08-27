# ─── build stage ──────────────────────────────────────────────
FROM golang:1.22-alpine AS builder
WORKDIR /src

# Disable the default module proxy that’s causing TLS errors
ENV GOPROXY=direct \
    GOSUMDB=off

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/task-service ./cmd/taskservice