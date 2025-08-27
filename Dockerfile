# ─── build stage ──────────────────────────────────────────────
FROM golang:1.22-alpine AS builder

# Switch APK mirrors from https → http to avoid TLS MITM problems,
# then install git so 'go mod download' works.
RUN sed -i 's/https/http/g' /etc/apk/repositories && \
    apk add --no-cache git

# Bypass Go module proxy cert issues (optional)
ENV GOPROXY=direct \
    GOSUMDB=off

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/task-service ./cmd/taskservice