// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: queries.sql

package db

import (
	"context"
)

const listIncidentTypes = `-- name: ListIncidentTypes :many


SELECT id, name FROM incident_type
`

// noinspection SqlDialectInspectionForFile
// noinspection SqlNoDataSourceInspectionForFile
func (q *Queries) ListIncidentTypes(ctx context.Context) ([]IncidentType, error) {
	rows, err := q.db.QueryContext(ctx, listIncidentTypes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []IncidentType
	for rows.Next() {
		var i IncidentType
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listIncidents = `-- name: ListIncidents :many
SELECT id, latitude, longtitude, incident_type_id, created_at, updated_at FROM incident
`

func (q *Queries) ListIncidents(ctx context.Context) ([]Incident, error) {
	rows, err := q.db.QueryContext(ctx, listIncidents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Incident
	for rows.Next() {
		var i Incident
		if err := rows.Scan(
			&i.ID,
			&i.Latitude,
			&i.Longtitude,
			&i.IncidentTypeID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
