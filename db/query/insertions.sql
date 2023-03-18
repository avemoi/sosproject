-- noinspection SqlDialectInspectionForFile

-- noinspection SqlNoDataSourceInspectionForFile

-- name: CreateInsident :execresult
INSERT INTO incident (
    power_id,
    latitude,
    longtitude
) VALUES (?, ?, ?);