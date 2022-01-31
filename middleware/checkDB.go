package middleware

import (
	"net/http"

	"github.com/leslesnoa/go-twitter/db"
)

/* DBのステータスチェックするミドルウェア*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := db.CheckingConnection(); err != nil {
			http.Error(w, "Bad connection internal server error", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
