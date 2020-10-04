package httpx

import "net/http"

// HSTS instructs browser to change all http:// requests to https://
func HSTS(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		s := "max-age=63072000; includeSubDomains"
		w.Header().Add("Strict-Transport-Security", s)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
