package repository

import (
	"database/sql"
	"fmt"

	"github.com/muratovdias/technodom-test-task/internal/app/models"
)

type Clinet interface {
	Redirect(link string) (models.Link, error)
}

type ClientRepo struct {
	db *sql.DB
}

func NewClient(db *sql.DB) *ClientRepo {
	return &ClientRepo{
		db: db,
	}
}

func (c *ClientRepo) Redirect(link string) (models.Link, error) {
	var res models.Link
	row := c.db.QueryRow(`SELECT * FROM links WHERE active_link=$1`, link)
	err := row.Scan(&res.ID, &res.ActiveLink, &res.HistoryLink)
	if err != nil {
		return res, fmt.Errorf("repository: client: redirect: %w", err)
	}
	return res, nil
}
