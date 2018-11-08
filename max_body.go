package httpx // import "kkn.fi/httpx"

import "net/http"

// MaxBody implements the https://golang.org/pkg/net/http/#MaxBytesReader
// as middleware. MaxBody reads maximum of given n bytes from the request
// body. MaxBody is only applied to HTTP methods POST and PUT.
func MaxBody(n int64) Middleware {
	f := func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost, http.MethodPut:
				r.Body = http.MaxBytesReader(w, r.Body, n)
			}
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
	return f
}
