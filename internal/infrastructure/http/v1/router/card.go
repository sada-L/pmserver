package router

import (
	"github.com/gorilla/mux"
	"github.com/sada-L/pmserver/internal/infrastructure/http/v1/controller"
)

func NewCardRouter(uc *controller.CardController, r *mux.Router) {
	r.Handle("/cards/current", uc.GetCardsByUser()).Methods("GET")
}
