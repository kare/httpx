package httpx

import "net/http"

const (
	contentTypeHeader = "Content-Type"
	charsetUTF8       = ";charset=utf-8"
	contenTypeJSON    = "application/json" + charsetUTF8
	contenTypeHTML    = "text/html" + charsetUTF8
	contentTypeText   = "text/plain" + charsetUTF8
)

// ContentTypeJSON sets the response ’Content-Type’ HTTP header with value `application/json;charset=utf-8`.
func ContentTypeJSON(h http.Handler) http.Handler {
	return ContentType(contenTypeJSON, h)
}

// ContentTypeHTML sets the response ’Content-Type’ HTTP header with value `text/html;charset=utf-8`.
func ContentTypeHTML(h http.Handler) http.Handler {
	return ContentType(contenTypeHTML, h)
}

// ContentTypeText sets the response ’Content-Type’ HTTP header with value `text/plain;charset=utf-8`.
func ContentTypeText(h http.Handler) http.Handler {
	return ContentType(contentTypeText, h)
}

// ContentType sets the response `Content-Type` HTTP header with given contentType.
func ContentType(contentType string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentTypeHeader, contentType)
		h.ServeHTTP(w, r)
	})
}
