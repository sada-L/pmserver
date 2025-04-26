package middleware

import (
	"net/http"
	"strings"

	"github.com/sada-L/pmserver/internal/model"
	"github.com/sada-L/pmserver/pkg/utils"
)

func AuthenticateMwf(us model.UserService) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Vary", "Authorization")
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				utils.InvalidAuthTokenError(w)
				return
			}

			ss := strings.Split(authHeader, " ")
			if len(ss) < 2 {
				utils.InvalidAuthTokenError(w)
				return
			}

			token := ss[1]
			claims, err := utils.ParseUserToken(token)
			if err != nil {
				utils.InvalidAuthTokenError(w)
				return
			}

			email := claims["email"].(string)
			user, err := us.UserByEmail(r.Context(), email)
			if err != nil {
				utils.InternalError(w, err)
				return
			}

			r = utils.SetContextUser(r, user)
			r = utils.SetContextUserToken(r, token)
			h.ServeHTTP(w, r)
		})
	}
}
