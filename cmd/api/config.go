package main

import (
	"database/sql"
	db "github.com/avemoi/sosproject/db/sqlc"
	"sync"
)

type models struct {
	db *db.Queries
}

func NewRepo(db *db.Queries) *models {
	return &models{db: db}
}

type Config struct {
	DB           *sql.DB
	env          string
	Wait         *sync.WaitGroup
	Models       *models
	GoogleApiKey string
	DeddieAddr   string
}
