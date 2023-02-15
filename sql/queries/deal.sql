-- name: CreateDeal :exec

INSERT INTO deal (id, user_id, type, value, description, trade_for, location_lat, location_lng,
                  location_address, location_city, location_state, location_zip_code, urgency_type,
                  urgency_limit_date, photos, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateDeal :exec
update deal
set user_id            = ?,
    type               = ?,
    value              = ?,
    description        = ?,
    trade_for          = ?,
    location_lat       = ?,
    location_lng       = ?,
    location_address   = ?,
    location_city      = ?,
    location_state     = ?,
    location_zip_code  = ?,
    urgency_type       = ?,
    urgency_limit_date = ?,
    photos             = ?,
    updated_at         = ?
where id = ?;

-- name: FindByIDDeal :one
SELECT *
FROM deal
WHERE id = ?;