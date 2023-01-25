-- name: CreateUser :exec
INSERT INTO user (id, name, email, login, password, location_lat, location_lng, location_address,
                              location_city, location_state, location_zip_code, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateUser :exec
UPDATE user SET name = ?, email = ?, login = ?, password = ?, location_lat = ?, location_lng = ?, location_address = ?,
                location_city = ?, location_state = ?, location_zip_code = ?, updated_at = ?
WHERE id = ?;

-- name: FindByLogin :one
SELECT * FROM user WHERE login = ?;

-- name: FindByEmail :one
SELECT * FROM user WHERE email = ?;

-- name: FindByID :one
SELECT * FROM user WHERE id = ?;