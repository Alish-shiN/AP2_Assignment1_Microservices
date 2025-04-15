package order

import (
    "context"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
    Create(ctx context.Context, order Order) (primitive.ObjectID, error)
    GetByID(ctx context.Context, id primitive.ObjectID) (Order, error)
    UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error
    List(ctx context.Context) ([]Order, error)
}

type service struct {
    repo Repository
}

func NewService(r Repository) Service {
    return &service{repo: r}
}

func (s *service) Create(ctx context.Context, o Order) (primitive.ObjectID, error) {
    return s.repo.Create(ctx, o)
}

func (s *service) GetByID(ctx context.Context, id primitive.ObjectID) (Order, error) {
    return s.repo.GetByID(ctx, id)
}

func (s *service) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error {
    return s.repo.UpdateStatus(ctx, id, status)
}

func (s *service) List(ctx context.Context) ([]Order, error) {
    return s.repo.List(ctx)
}
