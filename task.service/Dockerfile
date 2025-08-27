# ─── build stage ──────────────────────────────────────────────
FROM golang:1.22-alpine AS builder

# (Optional) switch APK mirrors to HTTP to avoid TLS issues, then install git
RUN sed -i 's/https/http/g' /etc/apk/repositories && \
    apk add --no-cache git

WORKDIR /src

# Copy entire project, including the vendor directory you created on host
COPY . .

# Build using the local vendor folder – NO network traffic inside container
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -mod=vendor -o task-service ./cmd/server

# ─── runtime stage ────────────────────────────────────────────
FROM gcr.io/distroless/static
COPY --from=builder /src/task-service /task-service
USER 10001
EXPOSE 8080
ENTRYPOINT ["/task-service"]