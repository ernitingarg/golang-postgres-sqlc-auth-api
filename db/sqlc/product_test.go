package db

import (
	"context"
	"testing"
	"time"

	"github.com/ernitingarg/golang-postgres-sqlc-auth-api/utils"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func createRandomProduct(t *testing.T) Product {
	customer := createRandomCustomer(t)

	arg := CreateProductParams{
		Title:      utils.RandomString(3),
		Content:    pgtype.Text{String: utils.RandomString(10), Valid: true},
		Price:      float64(utils.RandomNumber(1, 3)),
		CustomerID: customer.ID,
	}

	product, err := testQueries.CreateProduct(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, product)

	assert.NotZero(t, product.ID)
	assert.Equal(t, arg.Title, product.Title)
	assert.Equal(t, arg.Content, product.Content)
	assert.Equal(t, arg.Price, product.Price)
	assert.Equal(t, arg.CustomerID, product.CustomerID)
	assert.NotEmpty(t, product.CreatedAt)
	assert.NotEmpty(t, product.UpdatedAt)

	return product
}

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}

func TestGetProduct(t *testing.T) {
	product1 := createRandomProduct(t)

	product2, err := testQueries.GetProduct(context.Background(), product1.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, product2)

	assert.Equal(t, product1.ID, product2.ID)
	assert.Equal(t, product1.Title, product2.Title)
	assert.Equal(t, product1.Content, product2.Content)
	assert.Equal(t, product1.Price, product2.Price)
	assert.Equal(t, product1.CustomerID, product2.CustomerID)
	assert.WithinDuration(t, product1.CreatedAt.Time, product2.CreatedAt.Time, time.Second)
	assert.WithinDuration(t, product1.UpdatedAt.Time, product2.UpdatedAt.Time, time.Second)
}

func TestUpdateProduct(t *testing.T) {
	product1 := createRandomProduct(t)

	arg := UpdateProductParams{
		ID:      product1.ID,
		Title:   utils.RandomString(5),
		Content: pgtype.Text{String: utils.RandomString(20), Valid: true},
		Price:   float64(utils.RandomNumber(5, 10)),
	}

	product2, err := testQueries.UpdateProduct(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, product2)

	assert.Equal(t, product1.ID, product2.ID)
	assert.Equal(t, arg.Title, product2.Title)
	assert.Equal(t, arg.Content, product2.Content)
	assert.Equal(t, arg.Price, product2.Price)
	assert.Equal(t, product1.CustomerID, product2.CustomerID)
	assert.WithinDuration(t, product1.CreatedAt.Time, product2.CreatedAt.Time, time.Second)
	assert.WithinDuration(t, product1.UpdatedAt.Time, product2.UpdatedAt.Time, time.Second)
}

func TestDeleteProduct(t *testing.T) {
	product1 := createRandomProduct(t)

	err := testQueries.DeleteProduct(context.Background(), product1.ID)
	assert.NoError(t, err)

	product2, err := testQueries.GetProduct(context.Background(), product1.ID)
	assert.Error(t, err)
	assert.EqualError(t, err, pgx.ErrNoRows.Error())
	assert.Empty(t, product2)
}

func TestListProducts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProduct(t)
	}

	arg := ListProductsParams{
		Limit:  5,
		Offset: 5,
	}

	products, err := testQueries.ListProducts(context.Background(), arg)
	assert.NoError(t, err)
	assert.Len(t, products, 5)

	for _, product := range products {
		assert.NotEmpty(t, product)
	}
}
