package httpx // import "kkn.fi/httpx"

import "net/http"

// MaxBytesReader implements the https://golang.org/pkg/net/http/#MaxBytesReader
// as middleware. MaxBytesReader reads maximum of given n bytes from the request
// body. MaxBytesReader is only applied to HTTP methods POST and PUT.
func MaxBytesReader(n int64) func(http.Handler) http.Handler {
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
