-- name: CreateUser :one
INSERT INTO users (
    email,
    password_hash,
    full_name,
    phone_number
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET 
    email = COALESCE($2, email),
    full_name = COALESCE($3, full_name),
    phone_number = COALESCE($4, phone_number),
    updated_at = NOW()
WHERE id = $1
RETURNING *;
