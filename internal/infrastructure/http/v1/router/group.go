package router

import (
	"github.com/gorilla/mux"
	"github.com/sada-L/pmserver/internal/infrastructure/http/v1/controller"
)

func NewGroupRouter(uc *controller.GroupController, r *mux.Router) {
	r.Handle("/groups", uc.GetGroupsByUser()).Methods("GET")
	r.Handle("/groups", uc.CreateGroup()).Methods("POST")
	r.Handle("/groups/{id}", uc.UpdateGroup()).Methods("PUT")
	r.Handle("/groups/{id}", uc.DeleteGroup()).Methods("DELETE")
}
