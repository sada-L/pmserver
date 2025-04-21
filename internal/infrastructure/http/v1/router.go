package v1

import (
	"os"

	"github.com/rs/cors"
	"github.com/sada-L/pmserver/config"
	"github.com/sada-L/pmserver/internal/infrastructure/http/middleware"
	"github.com/sada-L/pmserver/internal/infrastructure/http/v1/controller"
	"github.com/sada-L/pmserver/internal/infrastructure/http/v1/router"
	"github.com/sada-L/pmserver/internal/service"
	"github.com/sada-L/pmserver/pkg/postgres"
	"github.com/sada-L/pmserver/pkg/server"
)

func Setup(cfg *config.Config, db *postgres.DB, s *server.Server) {
	s.Router.Use(cors.AllowAll().Handler)
	s.Router.Use(middleware.Logger(os.Stdout))

	us := service.NewUserService(db)
	uc := controller.NewUserController(us)

	apiRouter := s.Router.PathPrefix("/api/v1").Subrouter()

	// puglic routes
	noAuth := apiRouter.PathPrefix("").Subrouter()
	router.NewHealthRouter(noAuth)
	router.NewPublicUserRouter(uc, noAuth)

	// optional routes
	// optinalAuth := apiRouter.PathPrefix("").Subrouter()
	// optionalAuth.Use(middleware.AuthenticateMwf(us))
	//
	// private routes
	authApiRoutes := apiRouter.PathPrefix("").Subrouter()
	authApiRoutes.Use(middleware.AuthenticateMwf(us))
	router.NewPrivateUserRouter(uc, authApiRoutes)
	//
	// router.NewUserRouter(uc, publicRouter)
	// router.NewSwaggerRouter(publicRouter)
	s.Server.Handler = s.Router
}
