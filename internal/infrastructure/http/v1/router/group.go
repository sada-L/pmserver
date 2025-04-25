package router

import (
	"github.com/gorilla/mux"
	"github.com/sada-L/pmserver/internal/infrastructure/http/v1/controller"
)

func NewGroupRouter(uc *controller.GroupController, r *mux.Router) {
	r.Handle("/groups/current", uc.GetGroupsByUser()).Methods("GET")
}
