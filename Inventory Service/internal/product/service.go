package product

import (
    "context"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
    Create(ctx context.Context, p Product) (primitive.ObjectID, error)
    GetByID(ctx context.Context, id primitive.ObjectID) (Product, error)
    Update(ctx context.Context, id primitive.ObjectID, p Product) error
    Delete(ctx context.Context, id primitive.ObjectID) error
    List(ctx context.Context) ([]Product, error)
}

type service struct {
    repo Repository
}

func NewService(r Repository) Service {
    return &service{repo: r}
}

func (s *service) Create(ctx context.Context, p Product) (primitive.ObjectID, error) {
    return s.repo.Create(ctx, p)
}

func (s *service) GetByID(ctx context.Context, id primitive.ObjectID) (Product, error) {
    return s.repo.GetByID(ctx, id)
}

func (s *service) Update(ctx context.Context, id primitive.ObjectID, p Product) error {
    return s.repo.Update(ctx, id, p)
}

func (s *service) Delete(ctx context.Context, id primitive.ObjectID) error {
    return s.repo.Delete(ctx, id)
}

func (s *service) List(ctx context.Context) ([]Product, error) {
    return s.repo.List(ctx)
}
