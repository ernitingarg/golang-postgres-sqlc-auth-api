-- name: GetCustomer :one
SELECT * FROM customers
WHERE id = $1 LIMIT 1;

-- name: ListCustomers :many
SELECT * FROM customers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateCustomer :one
INSERT INTO customers (
  email,
  name,
  hash_password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateCustomer :one
UPDATE customers
  set name = $2,
  hash_password = $3
WHERE id = $1
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE id = $1;