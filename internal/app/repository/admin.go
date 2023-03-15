package repository

import (
	"database/sql"
	"fmt"

	"github.com/muratovdias/technodom-test-task/internal/app/models"
)

const (
	selectAllLinks   = `SELECT * FROM links LIMIT 25 OFFSET $2;`
	selectLink       = `SELECT * FROM links WHERE id=$1;`
	selectActiveLink = `SELECT active_link FROM links WHERE id=$1;`
	updateActiveLink = `UPDATE links SET active_link=$1, history_link=$2 WHERE id=$3;`
	deleteLink       = `DELETE FROM links WHERE id=$1;`
	createLink       = `INSERT INTO links (active_link) VALUES ($1);`
)

type Admin interface {
	GetLinks(offset int) (*[]models.Link, error)
	GetLinkByID(id int) (models.Link, error)
	UpdateLink(id int, newActiveLink string) error
	DeleteLink(id int) error
	CreateLink(newLink string) error
}

type AdminRepo struct {
	db *sql.DB
}

func NewAdmin(db *sql.DB) *AdminRepo {
	return &AdminRepo{
		db: db,
	}
}

func (a *AdminRepo) GetLinks(offset int) (*[]models.Link, error) {
	rows, err := a.db.Query(selectAllLinks, offset)
	if err != nil {
		return nil, fmt.Errorf("repository: select all links: %w", err)
	}
	links := make([]models.Link, 0, 25)
	for rows.Next() {
		var link models.Link
		err := rows.Scan(&link.ID, &link.ActiveLink, &link.HistoryLink)
		if err != nil {
			return nil, fmt.Errorf("repository: scaning links: %w", err)
		}
		links = append(links, link)
	}
	return &links, nil
}

func (a *AdminRepo) GetLinkByID(id int) (models.Link, error) {
	var link models.Link
	err := a.db.QueryRow(selectLink, id).Scan(&link.ID, &link.ActiveLink, &link.HistoryLink)
	if err != nil {
		return link, fmt.Errorf("repository: select link by id: %w", err)
	}
	return link, nil
}

func (a *AdminRepo) UpdateLink(id int, newActiveLink string) error {
	var newHistoryLink string
	err := a.db.QueryRow(selectActiveLink, id).Scan(&newHistoryLink)
	if err != nil {
		return fmt.Errorf("repository: select active_link by id: %w", err)
	}
	stmt, err := a.db.Prepare(updateActiveLink)
	if err != nil {
		return fmt.Errorf("repository: prepare update active_link by id: %w", err)
	}
	_, err = stmt.Exec(newActiveLink, newHistoryLink, id)
	if err != nil {
		return fmt.Errorf("repository: exec update active_link by id: %w", err)
	}
	return nil
}

func (a *AdminRepo) DeleteLink(id int) error {
	stmt, err := a.db.Prepare(deleteLink)
	if err != nil {
		return fmt.Errorf("repository: prepare delete row by id: %w", err)
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("repository: exec delete row by id: %w", err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository: delete rows affected : %w", err)
	}
	if n == 0 {
		return fmt.Errorf("repository: %w", sql.ErrNoRows)
	}
	return nil
}

func (a *AdminRepo) CreateLink(newLink string) error {
	stmt, err := a.db.Prepare(createLink)
	if err != nil {
		return fmt.Errorf("repository: prepare delete row by id: %w", err)
	}
	_, err = stmt.Exec(newLink)
	if err != nil {
		return fmt.Errorf("repository: exec delete row by id: %w", err)
	}
	return nil
}
