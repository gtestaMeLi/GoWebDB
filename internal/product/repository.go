package product

import (
	"context"
	"database/sql"

	"github.com/gtestaMeLi/GoWebDB/internal/domain"
)

type Repository interface {
	GetByName(ctx context.Context, name string) (domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByName(ctx context.Context, name string) (domain.Product, error) {
	return domain.Product{}, nil
}
