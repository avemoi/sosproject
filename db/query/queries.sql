-- noinspection SqlDialectInspectionForFile

-- noinspection SqlNoDataSourceInspectionForFile


-- name: ListIncidents :many
SELECT * FROM incident;

-- name: ListUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT power_id, latitude, longitude
from users where id = ? LIMIT 1;

-- name: GetUserByPowerId :one
SELECT id, latitude, longitude
from users
where power_id = ? LIMIT 1;

-- name: GetNumberOfIncidents :many
SELECT id, user_id FROM incident WHERE created_at >= NOW() - INTERVAL ? MINUTE;