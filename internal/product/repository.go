package product

import (
	"context"
	"database/sql"

	"github.com/gtestaMeLi/GoWebDB/internal/domain"
)

type Repository interface {
	GetByName(ctx context.Context, name string) (domain.Product, error)
	Save(ctx context.Context, p domain.Product) (int, error)
	Get(ctx context.Context, id int) (domain.Product, error)
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

func (r *repository) Get(ctx context.Context, id int) (domain.Product, error) {
	query := "SELECT * FROM products WHERE id=?;"
	row := r.db.QueryRow(query, id)
	p := domain.Product{}
	err := row.Scan(&p.ID, &p.Name, &p.Type, &p.Count, &p.Price)
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}

func (r *repository) Save(ctx context.Context, p domain.Product) (int, error) {
	query := "INSERT INTO products (name,type,count,price) VALUES (?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&p.Name, &p.Type, &p.Count, &p.Price)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
