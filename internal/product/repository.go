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
	query := "SELECT * FROM products WHERE name = ?;"
	row := r.db.QueryRow(query, name)
	p := domain.Product{}
	err := row.Scan(&p.ID, &p.Name, &p.Type, &p.Count, &p.Price)
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}
