# ─── build stage ──────────────────────────────────────────────────────────────
FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/task-service ./cmd/taskservice

# ─── runtime stage ────────────────────────────────────────────────────────────
FROM gcr.io/distroless/static
COPY --from=builder /bin/task-service /task-service
USER 10001
EXPOSE 8080
ENTRYPOINT ["/task-service"]
