package main

import (
    "context"
    "log"
    "time"

    "task-service/internal/config"
    "task-service/internal/handler"
    mongorepo "task-service/internal/repository/mongo"
    "task-service/internal/service"
    "task-service/internal/transport"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    cfg := config.Load()

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
    if err != nil {
        log.Fatal(err)
    }

    col := client.Database(cfg.DBName).Collection("tasks")

    repo := mongorepo.NewTaskRepository(col)
    svc := service.NewTaskService(repo)
    th := handler.NewTaskHandler(svc)
    server := transport.NewServer(th)

    server.Start(":" + cfg.Port)
}
