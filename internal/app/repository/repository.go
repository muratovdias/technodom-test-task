package repository

import "database/sql"

type Repository struct {
	Admin
	Clinet
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Admin:  NewAdmin(db),
		Clinet: NewClient(db),
	}
}
