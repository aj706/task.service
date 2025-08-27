package service

import (
    "context"
    "time"
    "user-service/internal/model"
    "user-service/internal/repository"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct { repo repository.UserRepository }

func New(r repository.UserRepository) *UserService { return &UserService{r} }

func (s *UserService) Create(ctx context.Context, u *model.User) error {
    now := time.Now(); u.ID = primitive.NewObjectID(); u.CreatedAt = now; u.UpdatedAt = now
    return s.repo.Create(ctx, u)
}

func (s *UserService) Get(ctx context.Context, id primitive.ObjectID) (*model.User, error) { return s.repo.Get(ctx, id) }
func (s *UserService) Update(ctx context.Context, u *model.User) error { u.UpdatedAt = time.Now(); return s.repo.Update(ctx, u) }
func (s *UserService) Delete(ctx context.Context, id primitive.ObjectID) error { return s.repo.Delete(ctx, id) }
