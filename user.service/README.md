# Task/User Micro-service Demo

See root README for full details. This folder contains **user.service**, a standalone Go micro-service that manages users in MongoDB.

## Quick Start

```bash
docker build -t user-service .
docker run -p 8081:8081 \
  -e MONGO_URI="mongodb+srv://<user>:<pass>@cluster.mongodb.net" \
  -e DB_NAME=users \
  -e PORT=8081 \
  user-service
```

### Endpoints

| Method | Path | Body | Notes |
|--------|------|------|-------|
| POST   | `/api/v1/users` | `{ "name":"...", "email":"..." }` | Create |
| GET    | `/api/v1/users/{id}` | — | Fetch one |
| PUT    | `/api/v1/users/{id}` | full doc | Replace |
| DELETE | `/api/v1/users/{id}` | — | Delete |
| GET    | `/healthz` | — | Liveness |
