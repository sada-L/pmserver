package router

import (
	"github.com/gorilla/mux"
	"github.com/sada-L/pmserver/internal/infrastructure/http/v1/controller"
)

func NewPublicUserRouter(uc *controller.UserController, r *mux.Router) {
	r.Handle("/users/login", uc.LoginUser()).Methods("POST")
	r.Handle("/users/create", uc.CreateUser()).Methods("POST")
}

func NewPrivateUserRouter(uc *controller.UserController, r *mux.Router) {
	r.Handle("/users/current", uc.GetCurrentUser()).Methods("GET")
}
