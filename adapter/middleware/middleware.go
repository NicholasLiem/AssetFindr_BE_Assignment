package middleware

import (
	"net/http"
)

func JSONMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		next(rw, r)
	}
}
