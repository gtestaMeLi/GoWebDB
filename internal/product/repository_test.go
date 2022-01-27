package product

import (
	"context"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gtestaMeLi/GoWebDB/internal/domain"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_sqlRepository_Store_Mock(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO products")
	mock.ExpectExec("INSERT INTO products").WillReturnResult(sqlmock.NewResult(5, 1))

	columns := []string{"id", "name", "type", "count", "price"}
	rows := sqlmock.NewRows(columns)
	productId := 5
	rows.AddRow(productId, "", "", 0, 0)

	mock.ExpectQuery("SELECT id, name, type, count, price FROM products WHERE id=?").WithArgs(productId).WillReturnRows(rows)

	repo := NewRepository(db)
	ctx := context.TODO()
	product := domain.Product{
		ID: productId,
	}
	_, err = repo.Save(ctx, product)
	assert.NoError(t, err)
	getResult, err := repo.Get(ctx, productId)
	assert.NoError(t, err)
	assert.NotNil(t, getResult)
	assert.Equal(t, product.ID, getResult.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}
