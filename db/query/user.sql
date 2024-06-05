-- name: CreateUser :one

INSERT INTO users (
  id,
  hashed_password,
  full_name,
  email
)
VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserForUpdate :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1
FOR NO KEY UPDATE;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET 
  full_name = $2,
  email = $3
WHERE id = $1
RETURNING *;

-- name: UpdateUserPassword :one
UPDATE users
SET 
  hashed_password = $2,
  password_changed_at = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

