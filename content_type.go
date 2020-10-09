package httpx

import (
	"net/http"

	"kkn.fi/httpx/internal/contenttype"
)

// ContentTypeJSON sets the response ’Content-Type’ HTTP header with value `application/json;charset=utf-8`.
func ContentTypeJSON(h http.Handler) http.Handler {
	return ContentType(contenttype.JSON, h)
}

// ContentTypeHTML sets the response ’Content-Type’ HTTP header with value `text/html;charset=utf-8`.
func ContentTypeHTML(h http.Handler) http.Handler {
	return ContentType(contenttype.HTML, h)
}

// ContentTypeText sets the response ’Content-Type’ HTTP header with value `text/plain;charset=utf-8`.
func ContentTypeText(h http.Handler) http.Handler {
	return ContentType(contenttype.Text, h)
}

// ContentType sets the response `Content-Type` HTTP header with given ct (HTML, JSON, etc).
func ContentType(ct contenttype.Type, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		value := ct.String() + contenttype.CharsetUTF8
		w.Header().Set(contenttype.Header, value)
		h.ServeHTTP(w, r)
	})
}
