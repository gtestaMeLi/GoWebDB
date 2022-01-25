package product

import (
	"context"

	"github.com/gtestaMeLi/GoWebDB/internal/domain"
)

type Service interface {
	GetByName(ctx context.Context, name string) (domain.Product, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) GetByName(ctx context.Context, name string) (domain.Product, error) {
	return domain.Product{}, nil
}
