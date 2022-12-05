package middlewares

import (
	"net/http"
)

func UserValidation(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("user_id")
		uuid := r.URL.Query().Get("user_uuid")

		if id == "" || uuid == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

