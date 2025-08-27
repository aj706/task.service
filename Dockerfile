# ─── build stage ──────────────────────────────────────────────
FROM golang:1.22-alpine AS builder

# Install git so go can download modules
RUN apk add --no-cache git

# Bypass corporate proxy TLS issue (optional – keep if you still get x509 errors)
ENV GOPROXY=direct \
    GOSUMDB=off

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/task-service ./cmd/taskservice