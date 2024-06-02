// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createLoan = `-- name: CreateLoan :one
INSERT INTO loans (user_id, amount) VALUES ($1, $2) RETURNING id, user_id, amount, created_at, updated_at
`

type CreateLoanParams struct {
	UserID uuid.UUID      `db:"user_id" json:"user_id"`
	Amount pgtype.Numeric `db:"amount" json:"amount"`
}

func (q *Queries) CreateLoan(ctx context.Context, arg CreateLoanParams) (Loan, error) {
	row := q.db.QueryRow(ctx, createLoan, arg.UserID, arg.Amount)
	var i Loan
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Amount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createRole = `-- name: CreateRole :one
INSERT INTO roles (name) VALUES ($1) RETURNING id, name, created_at, updated_at
`

func (q *Queries) CreateRole(ctx context.Context, name string) (Role, error) {
	row := q.db.QueryRow(ctx, createRole, name)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email, password, created_at, updated_at
`

type CreateUserParams struct {
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Name, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUserRole = `-- name: CreateUserRole :one
INSERT INTO mapping_users_roles (user_id, role_id) VALUES ($1, $2) RETURNING user_id, role_id
`

type CreateUserRoleParams struct {
	UserID uuid.UUID `db:"user_id" json:"user_id"`
	RoleID uuid.UUID `db:"role_id" json:"role_id"`
}

func (q *Queries) CreateUserRole(ctx context.Context, arg CreateUserRoleParams) (MappingUsersRole, error) {
	row := q.db.QueryRow(ctx, createUserRole, arg.UserID, arg.RoleID)
	var i MappingUsersRole
	err := row.Scan(&i.UserID, &i.RoleID)
	return i, err
}

const deleteLoan = `-- name: DeleteLoan :one
DELETE FROM loans WHERE id = $1 RETURNING id, user_id, amount, created_at, updated_at
`

func (q *Queries) DeleteLoan(ctx context.Context, id uuid.UUID) (Loan, error) {
	row := q.db.QueryRow(ctx, deleteLoan, id)
	var i Loan
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Amount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteRole = `-- name: DeleteRole :one
DELETE FROM roles WHERE id = $1 RETURNING id, name, created_at, updated_at
`

func (q *Queries) DeleteRole(ctx context.Context, id uuid.UUID) (Role, error) {
	row := q.db.QueryRow(ctx, deleteRole, id)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE FROM users WHERE id = $1 RETURNING id, name, email, password, created_at, updated_at
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, deleteUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUsersRoles = `-- name: DeleteUsersRoles :one
DELETE FROM mapping_users_roles WHERE user_id = $1 AND role_id = $2 RETURNING user_id, role_id
`

type DeleteUsersRolesParams struct {
	UserID uuid.UUID `db:"user_id" json:"user_id"`
	RoleID uuid.UUID `db:"role_id" json:"role_id"`
}

func (q *Queries) DeleteUsersRoles(ctx context.Context, arg DeleteUsersRolesParams) (MappingUsersRole, error) {
	row := q.db.QueryRow(ctx, deleteUsersRoles, arg.UserID, arg.RoleID)
	var i MappingUsersRole
	err := row.Scan(&i.UserID, &i.RoleID)
	return i, err
}

const getLoanById = `-- name: GetLoanById :one
SELECT id, user_id, amount, created_at, updated_at FROM loans WHERE id = $1
`

func (q *Queries) GetLoanById(ctx context.Context, id uuid.UUID) (Loan, error) {
	row := q.db.QueryRow(ctx, getLoanById, id)
	var i Loan
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Amount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getLoans = `-- name: GetLoans :many
SELECT id, user_id, amount, created_at, updated_at FROM loans
`

func (q *Queries) GetLoans(ctx context.Context) ([]Loan, error) {
	rows, err := q.db.Query(ctx, getLoans)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Loan
	for rows.Next() {
		var i Loan
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Amount,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRoleById = `-- name: GetRoleById :one
SELECT id, name, created_at, updated_at FROM roles WHERE id = $1
`

func (q *Queries) GetRoleById(ctx context.Context, id uuid.UUID) (Role, error) {
	row := q.db.QueryRow(ctx, getRoleById, id)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRoles = `-- name: GetRoles :many
SELECT id, name, created_at, updated_at FROM roles
`

func (q *Queries) GetRoles(ctx context.Context) ([]Role, error) {
	rows, err := q.db.Query(ctx, getRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Role
	for rows.Next() {
		var i Role
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserById = `-- name: GetUserById :one
SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUserEmail = `-- name: GetUserByUserEmail :one
SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByUserEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUserEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, name, email, password, created_at, updated_at FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsersRolesByRoleId = `-- name: GetUsersRolesByRoleId :many
SELECT user_id, role_id FROM mapping_users_roles WHERE role_id = $1
`

func (q *Queries) GetUsersRolesByRoleId(ctx context.Context, roleID uuid.UUID) ([]MappingUsersRole, error) {
	rows, err := q.db.Query(ctx, getUsersRolesByRoleId, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MappingUsersRole
	for rows.Next() {
		var i MappingUsersRole
		if err := rows.Scan(&i.UserID, &i.RoleID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsersRolesByUserId = `-- name: GetUsersRolesByUserId :many
SELECT user_id, role_id FROM mapping_users_roles WHERE user_id = $1
`

func (q *Queries) GetUsersRolesByUserId(ctx context.Context, userID uuid.UUID) ([]MappingUsersRole, error) {
	rows, err := q.db.Query(ctx, getUsersRolesByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MappingUsersRole
	for rows.Next() {
		var i MappingUsersRole
		if err := rows.Scan(&i.UserID, &i.RoleID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateLoan = `-- name: UpdateLoan :one
UPDATE loans SET user_id = $1, amount = $2 WHERE id = $3 RETURNING id, user_id, amount, created_at, updated_at
`

type UpdateLoanParams struct {
	UserID uuid.UUID      `db:"user_id" json:"user_id"`
	Amount pgtype.Numeric `db:"amount" json:"amount"`
	ID     uuid.UUID      `db:"id" json:"id"`
}

func (q *Queries) UpdateLoan(ctx context.Context, arg UpdateLoanParams) (Loan, error) {
	row := q.db.QueryRow(ctx, updateLoan, arg.UserID, arg.Amount, arg.ID)
	var i Loan
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Amount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateRole = `-- name: UpdateRole :one
UPDATE roles SET name = $1 WHERE id = $2 RETURNING id, name, created_at, updated_at
`

type UpdateRoleParams struct {
	Name string    `db:"name" json:"name"`
	ID   uuid.UUID `db:"id" json:"id"`
}

func (q *Queries) UpdateRole(ctx context.Context, arg UpdateRoleParams) (Role, error) {
	row := q.db.QueryRow(ctx, updateRole, arg.Name, arg.ID)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4 RETURNING id, name, email, password, created_at, updated_at
`

type UpdateUserParams struct {
	Name     string    `db:"name" json:"name"`
	Email    string    `db:"email" json:"email"`
	Password string    `db:"password" json:"-"`
	ID       uuid.UUID `db:"id" json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}