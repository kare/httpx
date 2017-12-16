package httpx // import "kkn.fi/httpx"

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server wraps an http.Server and provides helper functions for serving and
// shutdown.
type Server struct {
	srv *http.Server
}

// NewServer creates a new HTTP server that listens on given addr and the
// underlying http.Server can configured through given options.
// If addr is empty then ":http" is used.
func NewServer(addr string, options ...func(*http.Server)) *Server {
	if addr == "" {
		addr = ":http"
	}
	srv := http.Server{
		Addr:              addr,
		ReadTimeout:       time.Second * 15,
		ReadHeaderTimeout: time.Second * 15,
		WriteTimeout:      time.Second * 15,
		IdleTimeout:       time.Second * 30,
		MaxHeaderBytes:    4096,
	}
	for _, option := range options {
		option(&srv)
	}
	return &Server{
		srv: &srv,
	}
}

// ListenAndServe listens on the TCP network address s.srv.Addr and then calls
// Serve to handle requests on incoming connections
func (s *Server) ListenAndServe() {
	log.Print("http server started")
	if err := s.srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// ListenAndServeTLS listens on the TCP network address s.srv.Addr and then calls
// ServeTLS to handle requests on incoming TLS connections.
func (s *Server) ListenAndServeTLS(certFile, keyFile string) {
	log.Print("https server started")
	if err := s.srv.ListenAndServeTLS(certFile, keyFile); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// Shutdown gracefully shuts down the server without interrupting any active
// connections.
func (s *Server) Shutdown() {
	s.srv.SetKeepAlivesEnabled(false)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Printf("server shutdown error: %v\n", err)
	} else {
		log.Print("server shutdown complete")
	}
}
