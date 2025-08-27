package mongo

import (
    "context"
    "errors"
    "user-service/internal/model"
    repo "user-service/internal/repository"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct { col *mongo.Collection }

func New(col *mongo.Collection) repo.UserRepository { return &userRepo{col} }

func (r *userRepo) Create(ctx context.Context, u *model.User) error { _, err := r.col.InsertOne(ctx, u); return err }

func (r *userRepo) Get(ctx context.Context, id primitive.ObjectID) (*model.User, error) {
    var u model.User
    err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&u)
    if err == mongo.ErrNoDocuments { return nil, errors.New("user not found") }
    return &u, err
}

func (r *userRepo) Update(ctx context.Context, u *model.User) error { _, err := r.col.ReplaceOne(ctx, bson.M{"_id": u.ID}, u); return err }

func (r *userRepo) Delete(ctx context.Context, id primitive.ObjectID) error { _, err := r.col.DeleteOne(ctx, bson.M{"_id": id}); return err }
