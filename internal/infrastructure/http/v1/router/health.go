package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sada-L/pmserver/pkg/utils"
)

func NewHealthRouter(r *mux.Router) {
	r.Handle("/health", healthCheck())
}

func healthCheck() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		resp := utils.M{
			"status":  "available",
			"message": "healthy",
			"data":    utils.M{"hello": "beautiful"},
		}
		utils.WriteJSON(rw, http.StatusOK, resp)
	})
}
