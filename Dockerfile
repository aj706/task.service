# ─── build stage ──────────────────────────────────────────────
FROM golang:1.22-alpine AS builder

# Install git so Go can fetch modules
RUN sed -i 's/https/http/g' /etc/apk/repositories && \
    apk add --no-cache git

# Bypass corporate proxy TLS issues for Go module downloads
ENV GOPROXY=direct \
    GOINSECURE=github.com,*.github.com \
    GOSUMDB=off \
    GIT_SSL_NO_VERIFY=1

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/task-service ./cmd/taskservice

# ─── runtime stage ────────────────────────────────────────────
FROM gcr.io/distroless/static
COPY --from=builder /bin/task-service /task-service
USER 10001
EXPOSE 8080
ENTRYPOINT ["/task-service"]