package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgres(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - sql.Open: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - db.Ping: %w", err)
	}

	log.Print("Connect to datebase successes")
	return db, nil
}
