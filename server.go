package httpx

import (
	"net/http"
	"time"
)

// NewServer creates a pre-configured [http.Server] with reasonable defaults.
func NewServer(addr string, handler http.Handler, options ...func(*http.Server)) (*http.Server, error) {
	s := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       30 * time.Second,
		MaxHeaderBytes:    8 * 1024,
	}
	for _, option := range options {
		if option != nil {
			option(s)
		}
	}
	return s, nil
}
