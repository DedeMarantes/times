package middleware

import "net/http"

func SetHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			h.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}
