package httpx

import (
	"context"
	"net/http"
)

// HandlerWithContext is similar to http.Handler, adds support to
// context.Context and centralized error handling.
type HandlerWithContext interface {
	ServeHTTPWithContext(context.Context, http.ResponseWriter, *http.Request) error
}

// HandlerWithContextFunc adapts any function to HandlerWithContext type.
type HandlerWithContextFunc func(context.Context, http.ResponseWriter, *http.Request) error

// ServeHTTPWithContext calls itself as a function.
func (h HandlerWithContextFunc) ServeHTTPWithContext(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return h(ctx, w, r)
}
