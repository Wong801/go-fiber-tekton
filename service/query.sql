-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateUser :one
UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4 RETURNING *;

-- name: DeleteUser :one
DELETE FROM users WHERE id = $1 RETURNING *;

-- name: GetRoles :many
SELECT * FROM roles;

-- name: GetRoleById :one
SELECT * FROM roles WHERE id = $1;

-- name: CreateRole :one
INSERT INTO roles (name) VALUES ($1) RETURNING *;

-- name: UpdateRole :one
UPDATE roles SET name = $1 WHERE id = $2 RETURNING *;

-- name: DeleteRole :one
DELETE FROM roles WHERE id = $1 RETURNING *;

-- name: GetUsersRolesByUserId :many
SELECT * FROM mapping_users_roles WHERE user_id = $1;

-- name: GetUserByUserEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUsersRolesByRoleId :many
SELECT * FROM mapping_users_roles WHERE role_id = $1;

-- name: CreateUserRole :one
INSERT INTO mapping_users_roles (user_id, role_id) VALUES ($1, $2) RETURNING *;

-- name: DeleteUsersRoles :one
DELETE FROM mapping_users_roles WHERE user_id = $1 AND role_id = $2 RETURNING *;

-- name: GetLoans :many
SELECT * FROM loans;

-- name: GetLoanById :one
SELECT * FROM loans WHERE id = $1;

-- name: CreateLoan :one
INSERT INTO loans (user_id, amount) VALUES ($1, $2) RETURNING *;

-- name: UpdateLoan :one
UPDATE loans SET user_id = $1, amount = $2 WHERE id = $3 RETURNING *;

-- name: DeleteLoan :one
DELETE FROM loans WHERE id = $1 RETURNING *;