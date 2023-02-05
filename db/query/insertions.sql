-- noinspection SqlDialectInspectionForFile

-- noinspection SqlNoDataSourceInspectionForFile

-- name: CreateInsidentType :execresult
INSERT INTO incident_type (
    name
) VALUES (?);

-- name: CreateInsident :execresult
INSERT INTO incident (
    longtitude,
    latitude,
    incident_type_id

) VALUES (?, ?, ?);