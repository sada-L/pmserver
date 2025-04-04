package v1

import (
	"database/sql"

	"github.com/sada-L/pmserver/config"
	_ "github.com/sada-L/pmserver/docs"
	"github.com/sada-L/pmserver/internal/infrastructure/http/v1/controller"
	"github.com/sada-L/pmserver/internal/infrastructure/http/v1/router"
	"github.com/sada-L/pmserver/internal/repository"
	"github.com/sada-L/pmserver/internal/service"
	"github.com/sada-L/pmserver/pkg/httpserver"
)

// NewRouter -.
// Swagger spec:
//
//	@title			Go Clean Template API
//	@description	Using a translation service as an example
//	@version		1.0
//	@host			localhost:8080
//	@BasePath		/v1
func Setup(cfg *config.Config, db *sql.DB, s *httpserver.Server) {
	uc := controller.NewUserController(service.NewUserService(repository.NewUserRepository(db)))

	publicRouter := s.App.Group("/v1")
	router.NewUserRouter(uc, publicRouter)
	router.NewSwaggerRouter(publicRouter)
}
