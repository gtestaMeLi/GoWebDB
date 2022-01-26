package product

import (
	"context"
	"fmt"

	"github.com/gtestaMeLi/GoWebDB/internal/domain"
)

type Service interface {
	GetByName(ctx context.Context, name string) (domain.Product, error)
	Store(ctx context.Context, name string, ptype string, count int, price float64) (domain.Product, error)
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

func (s *service) Store(ctx context.Context, name string, ptype string, count int, price float64) (domain.Product, error) {
	prod := domain.Product{
		Name:  name,
		Type:  ptype,
		Count: count,
		Price: price,
	}

	newId, err := s.repo.Save(ctx, prod)
	if err != nil {
		return domain.Product{}, err
	}

	newProd, err := s.repo.Get(ctx, newId)

	if err != nil {
		return domain.Product{}, err
	}

	return newProd, nil
}
