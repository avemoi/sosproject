-- noinspection SqlDialectInspectionForFile

-- noinspection SqlNoDataSourceInspectionForFile

-- name: CreateInsident :execresult
INSERT INTO incident (
    user_id
) VALUES (?);


-- name: CreateUser :execresult
INSERT into users (
    power_id,
    latitude,
    longitude
) VALUES (?, ?, ?)