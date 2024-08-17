package middleware

import "net/http"

func SetHeader(key, value string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(key, value)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
