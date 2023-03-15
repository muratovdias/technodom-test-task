package repository

import "database/sql"

type Clinet interface {
}

type ClientRepo struct {
	db *sql.DB
}

func NewClient(db *sql.DB) *ClientRepo {
	return &ClientRepo{
		db: db,
	}
}
