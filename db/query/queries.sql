-- noinspection SqlDialectInspectionForFile

-- noinspection SqlNoDataSourceInspectionForFile


-- name: ListIncidents :many
SELECT * FROM incident;

-- name: ListUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT power_id, latitude, longtitude
from users where id = ? LIMIT 1;