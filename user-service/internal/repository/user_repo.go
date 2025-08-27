package repository

import (
    "context"
    "user-service/internal/model"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
    Create(ctx context.Context, u *model.User) error
    Get(ctx context.Context, id primitive.ObjectID) (*model.User, error)
    Update(ctx context.Context, u *model.User) error
    Delete(ctx context.Context, id primitive.ObjectID) error
}
