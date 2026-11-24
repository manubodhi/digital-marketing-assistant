-- name: CreateUser :one
INSERT INTO users (
    id, email, password_hash, full_name, is_email_verified
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE LOWER(email) = LOWER($1) LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: VerifyUserEmail :exec
UPDATE users
SET is_email_verified = TRUE, updated_at = NOW()
WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE users
SET password_hash = $2, updated_at = NOW()
WHERE id = $1;
