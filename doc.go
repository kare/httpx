// Package httpx provides a convenience wrapper for http.Server and a middleware
// wrapper for http.MaxBytesReader.
//
// Server example:
//		options := func(srv *http.Server) {
//			srv.Handler = ...
//			srv.ReadTimeout = time.Second * 5
//		}
//		srv := httpx.NewServer(":http", options)
//		go srv.ListenAndServe()
//		srv.Shutdown()
//
// MaxBytesReader example for chi:
//		r := chi.NewRouter()
//		r.Use(httpx.MaxBytesReader(1 << 20))
//
package httpx // import "kkn.fi/httpx"
