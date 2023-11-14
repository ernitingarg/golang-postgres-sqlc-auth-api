package db

import (
	"context"
	"testing"
	"time"

	"github.com/ernitingarg/golang-postgres-sqlc-auth-api/utils"
	"github.com/jackc/pgx"
	"github.com/stretchr/testify/assert"
)

func createRandomCustomer(t *testing.T) Customer {
	hashedPassword, err := utils.HashPassword(utils.RandomString(8))
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	arg := CreateCustomerParams{
		Name:         utils.RandomString(6),
		Email:        utils.RandomEmail(),
		HashPassword: hashedPassword,
	}

	customer, err := testQueries.CreateCustomer(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, customer)
	assert.Equal(t, arg.Name, customer.Name)
	assert.Equal(t, arg.Email, customer.Email)
	assert.Equal(t, arg.HashPassword, customer.HashPassword)
	assert.NotZero(t, customer.ID)
	assert.NotEmpty(t, customer.CreatedAt)
	assert.NotEmpty(t, customer.UpdatedAt)

	return customer
}

func TestCreateCustomer(t *testing.T) {
	createRandomCustomer(t)
}

func TestGetCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)

	customer2, err := testQueries.GetCustomer(context.Background(), customer1.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, customer2)

	assert.Equal(t, customer1.ID, customer2.ID)
	assert.Equal(t, customer1.Name, customer2.Name)
	assert.Equal(t, customer1.Email, customer2.Email)
	assert.WithinDuration(t, customer1.CreatedAt.Time, customer2.CreatedAt.Time, time.Second)
	assert.WithinDuration(t, customer1.UpdatedAt.Time, customer2.UpdatedAt.Time, time.Second)
}

func TestUpdateCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)

	arg := UpdateCustomerParams{
		ID:   customer1.ID,
		Name: utils.RandomString(7),
	}

	customer2, err := testQueries.UpdateCustomer(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, customer2)

	assert.Equal(t, customer1.ID, customer2.ID)
	assert.Equal(t, arg.Name, customer2.Name)
	assert.Equal(t, customer1.Email, customer2.Email)
	assert.WithinDuration(t, customer1.CreatedAt.Time, customer2.CreatedAt.Time, time.Second)
	assert.WithinDuration(t, customer1.UpdatedAt.Time, customer2.UpdatedAt.Time, time.Second)
}

func TestDeleteCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)

	err := testQueries.DeleteCustomer(context.Background(), customer1.ID)
	assert.NoError(t, err)

	customer2, err := testQueries.GetCustomer(context.Background(), customer1.ID)
	assert.Error(t, err)
	assert.EqualError(t, err, pgx.ErrNoRows.Error())
	assert.Empty(t, customer2)
}

func TestListCustomers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCustomer(t)
	}

	arg := ListCustomersParams{
		Limit:  5,
		Offset: 5,
	}

	customers, err := testQueries.ListCustomers(context.Background(), arg)
	assert.NoError(t, err)
	assert.Len(t, customers, 5)

	for _, customer := range customers {
		assert.NotEmpty(t, customer)
	}
}
