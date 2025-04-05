package app

import (
	"fmt"
	"log"

	"github.com/sada-L/pmserver/config"
	v1 "github.com/sada-L/pmserver/internal/infrastructure/http/v1"
	"github.com/sada-L/pmserver/pkg/postgres"
	"github.com/sada-L/pmserver/pkg/server"
)

func Run(cfg *config.Config) {
	db, err := postgres.Open(cfg.DB.URL)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.Open: %w", err))
	}
	defer db.Close()

	srv := server.NewServer()
	v1.Setup(cfg, db, srv)

	log.Fatal(srv.Run(cfg.Http.Port))
}
