package app

import (
	"fmt"
	"log"

	"github.com/sada-L/pmserver/config"
	v1 "github.com/sada-L/pmserver/internal/infrastructure/http/v1"
	"github.com/sada-L/pmserver/pkg/httpserver"
	"github.com/sada-L/pmserver/pkg/postgres"
)

func Run(cfg *config.Config) {
	db, err := postgres.NewPostgres(cfg.DB.URL)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.NewPostgres: %w", err))
	}
	defer db.Close()

	s := httpserver.NewServer(cfg.Http.Port)
	v1.Setup(cfg, db, s)

	s.Run()
}
