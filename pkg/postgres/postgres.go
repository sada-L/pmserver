package postgres

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
}

func Open(url string) (*DB, error) {
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("postgres - Open - sqlx.Open: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("postgres - Open - db.Ping: %w", err)
	}

	log.Println("connect to datebase successes")
	return &DB{db}, nil
}
