package httpx

import "net/http"

const (
	contentTypeHeader = "Content-Type"
	charsetUTF8       = ";charset=utf-8"
)

// Type defines all supported Content-Type header values.
type Type string

func (t Type) String() string {
	return string(t)
}

const (
	// JSON content type.
	JSON Type = "application/json"
	// HTML content type.
	HTML Type = "text/html"
	// Text content type.
	Text Type = "text/plain"
)

// ContentTypeJSON sets the response ’Content-Type’ HTTP header with value `application/json;charset=utf-8`.
func ContentTypeJSON(h http.Handler) http.Handler {
	return ContentType(JSON, h)
}

// ContentTypeHTML sets the response ’Content-Type’ HTTP header with value `text/html;charset=utf-8`.
func ContentTypeHTML(h http.Handler) http.Handler {
	return ContentType(HTML, h)
}

// ContentTypeText sets the response ’Content-Type’ HTTP header with value `text/plain;charset=utf-8`.
func ContentTypeText(h http.Handler) http.Handler {
	return ContentType(Text, h)
}

// ContentType sets the response `Content-Type` HTTP header with given contentType.
func ContentType(contentType Type, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentTypeHeader, contentType.String()+charsetUTF8)
		h.ServeHTTP(w, r)
	})
}
