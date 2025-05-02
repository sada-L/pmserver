package router

import (
	"github.com/gorilla/mux"
	"github.com/sada-L/pmserver/internal/infrastructure/http/v1/controller"
)

func NewCardRouter(cc *controller.CardController, r *mux.Router) {
	r.Handle("/cards", cc.GetCardsByUser()).Methods("GET")
	r.Handle("/cards", cc.CreateCard()).Methods("POST")
	r.Handle("/cards/{id}", cc.UpdateCard()).Methods("PUT")
	r.Handle("/cards/{id}", cc.DeleteCard()).Methods("DELETE")
}
