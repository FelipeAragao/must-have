// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO user (id, name, email, login, password, location_lat, location_lng, location_address,
                              location_city, location_state, location_zip_code, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateUserParams struct {
	ID              string
	Name            string
	Email           string
	Login           string
	Password        string
	LocationLat     sql.NullFloat64
	LocationLng     sql.NullFloat64
	LocationAddress string
	LocationCity    string
	LocationState   string
	LocationZipCode int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Login,
		arg.Password,
		arg.LocationLat,
		arg.LocationLng,
		arg.LocationAddress,
		arg.LocationCity,
		arg.LocationState,
		arg.LocationZipCode,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const findByEmail = `-- name: FindByEmail :one
SELECT id, name, email, login, password, location_lat, location_lng, location_address, location_city, location_state, location_zip_code, created_at, updated_at FROM user WHERE email = ?
`

func (q *Queries) FindByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, findByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Login,
		&i.Password,
		&i.LocationLat,
		&i.LocationLng,
		&i.LocationAddress,
		&i.LocationCity,
		&i.LocationState,
		&i.LocationZipCode,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findByID = `-- name: FindByID :one
SELECT id, name, email, login, password, location_lat, location_lng, location_address, location_city, location_state, location_zip_code, created_at, updated_at FROM user WHERE id = ?
`

func (q *Queries) FindByID(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, findByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Login,
		&i.Password,
		&i.LocationLat,
		&i.LocationLng,
		&i.LocationAddress,
		&i.LocationCity,
		&i.LocationState,
		&i.LocationZipCode,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findByLogin = `-- name: FindByLogin :one
SELECT id, name, email, login, password, location_lat, location_lng, location_address, location_city, location_state, location_zip_code, created_at, updated_at FROM user WHERE login = ?
`

func (q *Queries) FindByLogin(ctx context.Context, login string) (User, error) {
	row := q.db.QueryRowContext(ctx, findByLogin, login)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Login,
		&i.Password,
		&i.LocationLat,
		&i.LocationLng,
		&i.LocationAddress,
		&i.LocationCity,
		&i.LocationState,
		&i.LocationZipCode,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE user SET name = ?, email = ?, login = ?, password = ?, location_lat = ?, location_lng = ?, location_address = ?,
                location_city = ?, location_state = ?, location_zip_code = ?, updated_at = ?
WHERE id = ?
`

type UpdateUserParams struct {
	Name            string
	Email           string
	Login           string
	Password        string
	LocationLat     sql.NullFloat64
	LocationLng     sql.NullFloat64
	LocationAddress string
	LocationCity    string
	LocationState   string
	LocationZipCode int64
	UpdatedAt       time.Time
	ID              string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.Name,
		arg.Email,
		arg.Login,
		arg.Password,
		arg.LocationLat,
		arg.LocationLng,
		arg.LocationAddress,
		arg.LocationCity,
		arg.LocationState,
		arg.LocationZipCode,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
