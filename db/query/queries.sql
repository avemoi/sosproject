-- noinspection SqlDialectInspectionForFile

-- noinspection SqlNoDataSourceInspectionForFile


-- name: ListIncidents :many
select inc.id as incident_id, usr.id as user_id, usr.latitude ,usr.longitude
from incident inc
         join users usr on usr.id = inc.user_id;

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
select inc.id as incident_id, usr.id as user_id, usr.latitude ,usr.longitude
from incident inc
         join users usr on usr.id = inc.user_id
where inc.created_at >= NOW() - INTERVAL ? MINUTE;

-- name: GetUserIncidents :many
select inc.id as incident_id, usr.id as user_id, usr.latitude ,usr.longitude
from incident inc
         join users usr on usr.id = inc.user_id
where inc.created_at >= NOW() - INTERVAL ? MINUTE and usr.id=?;

-- name: CountRecentIncidents :one
select COUNT(*)
from incident inc
where inc.created_at >= NOW() - INTERVAL ? MINUTE;