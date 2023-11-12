// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Customer struct {
	ID           int64              `json:"id"`
	Email        string             `json:"email"`
	Name         string             `json:"name"`
	HashPassword string             `json:"hashPassword"`
	CreatedAt    pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt    pgtype.Timestamptz `json:"updatedAt"`
}

type Product struct {
	ID         int32              `json:"id"`
	Title      string             `json:"title"`
	Content    pgtype.Text        `json:"content"`
	Price      float64            `json:"price"`
	CreatedAt  pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt  pgtype.Timestamptz `json:"updatedAt"`
	CustomerID int64              `json:"customerId"`
}
