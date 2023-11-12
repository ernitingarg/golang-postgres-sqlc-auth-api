-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateProduct :one
INSERT INTO products (
  title,
  content,
  price,
  customer_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateProduct :one
UPDATE products
  set title = $2,
  content = $3,
  price = $4
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;