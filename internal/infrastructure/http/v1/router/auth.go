package router

import (
	"github.com/gorilla/mux"
	"github.com/sada-L/pmserver/internal/infrastructure/http/v1/controller"
)

func NewAuthRouter(uc *controller.AuthController, r *mux.Router) {
	r.Handle("/auth/login", uc.Login()).Methods("POST")
	r.Handle("/auth/register", uc.Register()).Methods("POST")
	r.Handle("/auth/refresh", uc.RefreshToken()).Methods("POST")
}
