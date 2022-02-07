package middleware

import (
	"net/http"

	"github.com/leslesnoa/go-twitter/handlers"
)

func ValidJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := handlers.AccessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error invalid request token "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
