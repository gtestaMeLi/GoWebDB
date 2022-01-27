package product

import (
	"context"
	"database/sql"

	"github.com/gtestaMeLi/GoWebDB/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
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

const (
	queryGetAll    = "SELECT id, name, type, count, price FROM products;"
	queryGetByName = "SELECT id, name, type, count, price FROM products WHERE name = ?;"
	queryGetOne    = "SELECT id, name, type, count, price FROM products WHERE id=?;"
	querySave      = "INSERT INTO products (name,type,count,price) VALUES (?,?,?,?);"
)

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {

	rows, err := r.db.Query(queryGetAll)
	if err != nil {
		return nil, err
	}

	var prods []domain.Product

	for rows.Next() {
		p := domain.Product{}
		_ = rows.Scan(&p.ID, &p.Name, &p.Type, &p.Count, &p.Price)
		prods = append(prods, p)
	}

	return prods, nil
}

func (r *repository) GetByName(ctx context.Context, name string) (domain.Product, error) {
	row := r.db.QueryRow(queryGetByName, name)
	p := domain.Product{}
	err := row.Scan(&p.ID, &p.Name, &p.Type, &p.Count, &p.Price)
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Product, error) {
	row := r.db.QueryRow(queryGetOne, id)
	p := domain.Product{}
	err := row.Scan(&p.ID, &p.Name, &p.Type, &p.Count, &p.Price)
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}

func (r *repository) Save(ctx context.Context, p domain.Product) (int, error) {
	stmt, err := r.db.Prepare(querySave)
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
