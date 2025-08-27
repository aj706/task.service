# Task/User Micro-service Demo

This repository contains **two independent Go micro-services** that share a single MongoDB (Atlas) cluster.

| Service | Port | Purpose | Mongo Collection |
|---------|------|---------|------------------|
| **task.service** | 8080 | CRUD tasks (todo → in_progress → done) | `tasks` |
| **user.service** | 8081 | CRUD users | `users` |

Both are stateless containers built with multi-stage Dockerfiles. `docker-compose.yml` lets you run them together locally.

---

## 1. Problem Breakdown & Design Decisions

| Concern | Choice / Reasoning |
|---------|--------------------|
| Framework | Go + chi → minimal, idiomatic for micro-services |
| Persistence | MongoDB Atlas; each service owns its own collection |
| Architecture | Clean layers: handler → service → repository |
| Inter-service Comms | REST (Task → User) via `USER_SERVICE_BASE` env |
| Scalability | Stateless containers scaled horizontally; Mongo replica-set |
| Resilience | `/healthz` endpoints; Task validates `userId` on create/update |

---

## 2. Running the Services Locally

```bash
# build & start both services
docker compose build
docker compose up
```

Environment variables (already set in compose):
```
MONGO_URI  # Atlas connection string
DB_NAME    # tasks or users
PORT       # 8080 / 8081
USER_SERVICE_BASE=http://user-service:8081  # task.service only
```

---

## 3. API Reference

### User Service `http://localhost:8081/api/v1/users`

| Method | Path | Body | Notes |
|--------|------|------|-------|
| POST   | `/users` | `{ "name":"...", "email":"..." }` | Create user |
| GET    | `/users/{id}` | — | Fetch user |
| PUT    | `/users/{id}` | full user JSON | Replace |
| DELETE | `/users/{id}` | — | Delete |
| GET    | `/healthz` | — | Liveness |

### Task Service `http://localhost:8080/api/v1/tasks`

| Method | Path | Body | Notes |
|--------|------|------|-------|
| POST   | `/tasks` | `{ "title":"...", "description":"...", "userId":"ObjectID" }` | Validates userId |
| GET    | `/tasks` | — | `?limit&offset&status=todo|in_progress|done` |
| GET    | `/tasks/{id}` | — | Get one |
| PATCH  | `/tasks/{id}` | `{ "status":"done" }` | Partial update |
| DELETE | `/tasks/{id}` | — | Remove task |
| GET    | `/healthz` | — | Liveness |

---

## 4. Micro-service Concepts Demonstrated

1. **Single Responsibility** – each service owns one aggregate.
2. **Independent Deployability** – separate modules, Dockerfiles, images.
3. **API-level Integrity** – Task validates users via HTTP call.
4. **Horizontal Scaling** – stateless containers + shared DB.
5. **Observability** – health endpoints suitable for k8s liveness checks.
