// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"database/sql"
)

type Incident struct {
	ID        int64         `json:"id"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
	UserID    sql.NullInt64 `json:"user_id"`
}

type User struct {
	ID         int64        `json:"id"`
	PowerID    int32        `json:"power_id"`
	Latitude   float64      `json:"latitude"`
	Longtitude float64      `json:"longtitude"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
}
