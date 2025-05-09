package router

import (
	"github.com/gorilla/mux"
	"github.com/sada-L/pmserver/internal/infrastructure/http/v1/controller"
)

func NewAuthRouter(ac *controller.AuthController, r *mux.Router) {
	r.Handle("/auth/login", ac.Login()).Methods("POST")
	r.Handle("/auth/register", ac.Register()).Methods("POST")
	r.Handle("/auth/refresh", ac.RefreshToken()).Methods("POST")
}

func NewAuth2FaRouter(ac *controller.AuthController, r *mux.Router) {
	r.Handle("/auth/2fa/enable", ac.TwoFaEnable()).Methods("POST")
	r.Handle("/auth/2fa/verify", ac.TwoFaVerify()).Methods("POST")
	r.Handle("/auth/2fa/disable", ac.TwoFaDisable()).Methods("POST")
}
