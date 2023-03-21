// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: insertions.sql

package db

import (
	"context"
	"database/sql"
)

const createInsident = `-- name: CreateInsident :execresult


INSERT INTO incident (
    user_id
) VALUES (?)
`

// noinspection SqlDialectInspectionForFile
// noinspection SqlNoDataSourceInspectionForFile
func (q *Queries) CreateInsident(ctx context.Context, userID sql.NullInt64) (sql.Result, error) {
	return q.db.ExecContext(ctx, createInsident, userID)
}

const createUser = `-- name: CreateUser :execresult
INSERT into users (
    power_id,
    latitude,
    longtitude
) VALUES (?, ?, ?)
`

type CreateUserParams struct {
	PowerID    int32   `json:"power_id"`
	Latitude   float64 `json:"latitude"`
	Longtitude float64 `json:"longtitude"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser, arg.PowerID, arg.Latitude, arg.Longtitude)
}
