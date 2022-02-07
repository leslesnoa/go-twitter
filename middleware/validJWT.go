package middleware

import (
	"net/http"

	"github.com/leslesnoa/go-twitter/handlers"
)

func ValidJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_, _, _, err := handlers.AccessToken(r.Header.Get("Authorization"), ctx)
		if err != nil {
			http.Error(w, "Error invalid request token "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
