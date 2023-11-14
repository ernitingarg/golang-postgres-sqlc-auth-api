package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dataSourceName = "postgresql://admin:password123@localhost:5432/postgresdb?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	connPool, err := pgxpool.New(context.Background(), dataSourceName)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	defer connPool.Close()

	testQueries = New(connPool)

	os.Exit(m.Run())
}
