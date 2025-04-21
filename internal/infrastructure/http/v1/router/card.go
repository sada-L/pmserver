package router

import (
	"github.com/gorilla/mux"
	"github.com/sada-L/pmserver/internal/infrastructure/http/v1/controller"
)

func NewCardRouter(uc *controller.UserController, r *mux.Router) {
	r.Handle("/cards/", uc.LoginUser()).Methods("POST")
	r.Handle("/cards/create", uc.CreateUser()).Methods("POST")
}
