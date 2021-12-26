package middleware

import (
	"net/http"

	"github.com/leslesnoa/go-twitter/db"
)

/* DBのステータスチェックミドルウェア*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckingConnection() == 0 {
			http.Error(w, "Connection bad internal server error", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
