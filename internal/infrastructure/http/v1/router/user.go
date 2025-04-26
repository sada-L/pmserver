package router

import (
	"github.com/gorilla/mux"
	"github.com/sada-L/pmserver/internal/infrastructure/http/v1/controller"
)

func NewUserRouter(uc *controller.UserController, r *mux.Router) {
	r.Handle("/users/current", uc.GetCurrentUser()).Methods("GET")
}
