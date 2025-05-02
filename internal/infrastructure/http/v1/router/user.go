package router

import (
	"github.com/gorilla/mux"
	"github.com/sada-L/pmserver/internal/infrastructure/http/v1/controller"
)

func NewUserRouter(uc *controller.UserController, r *mux.Router) {
	r.Handle("/users/me", uc.GetCurrentUser()).Methods("GET")
	r.Handle("/users/me", uc.UpdateUser()).Methods("PUT")
	r.Handle("/users/me", uc.DeleteUser()).Methods("DELETE")
}
