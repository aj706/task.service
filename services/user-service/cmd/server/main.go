package main

import (
    "context"
    "time"
    "user-service/internal/config"
    "user-service/internal/handler"
    mongorepo "user-service/internal/repository/mongo"
    "user-service/internal/service"
    "user-service/internal/transport"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    cfg := config.Load()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second); defer cancel()
    client, _ := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
    col := client.Database(cfg.DBName).Collection("users")
    repo := mongorepo.New(col)
    svc := service.New(repo)
    h := handler.New(svc)
    srv := transport.New(h)
    srv.Start(":" + cfg.Port)
}
