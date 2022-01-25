package product

import (
	"context"
	"fmt"

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
	product, err := s.repo.GetByName(ctx, name)
	if product.ID == 0 {
		return domain.Product{}, fmt.Errorf("Error: no such product %s", name)
	}
	return product, err
}
