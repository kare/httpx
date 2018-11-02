package httpx

import "net/http"

type (
	// Middleware is an adapter for http.Handler.
	Middleware func(http.Handler) http.Handler

	// MiddlewareFunc is an adapter for http.HandlerFunc.
	MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc
)

// Chain chains middleware and executes chained middlewares in reversed order. For example:
// Given call Chain(hello, header, logger) the execution order is: 1. header 2. logger 3. hello
func Chain(f http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		f = middlewares[i](f)
	}
	return f
}

// ChainFunc chains middleware funcs and executes chained funcs in reversed order.
// Given call ChainFunc(hello, header, logger) the execution order is: 1. header 2. logger 3. hello
func ChainFunc(f http.HandlerFunc, middlewares ...MiddlewareFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		f = middlewares[i](f)
	}
	return f
}
