package main

import (
	"log"

	"github.com/sada-L/pmserver/config"
	"github.com/sada-L/pmserver/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
