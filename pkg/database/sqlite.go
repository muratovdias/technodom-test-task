package database

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/muratovdias/technodom-test-task/internal/app/models"
)

const schema = `
	CREATE TABLE IF NOT EXISTS links(
		id INTEGER PRIMARY KEY,
    	active_link TEXT UNIQUE NOT NULL,
    	history_link TEXT UNIQUE NOT NULL
	)
`

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InsertData(db *sql.DB) {
	file, err := os.ReadFile("links.json")
	if err != nil {
		log.Fatal(err)
	}
	var links []models.Link
	err = json.Unmarshal(file, &links)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO links(active_link, history_link) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	for _, link := range links {
		_, err = stmt.Exec(link.ActiveLink, link.HistoryLink)
		if err != nil {
			if strings.ContainsAny(err.Error(), "UNIQUE constraint failed") {
				return
			}
			tx.Rollback()
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
