package middleware

import "net/http"

var cacheControlHeader = http.CanonicalHeaderKey("Cache-Control")

func CacheControl(value string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(cacheControlHeader, value)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
