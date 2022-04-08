package httpx_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"kkn.fi/httpx"
)

func newRequest(t *testing.T, method, path, body string) *http.Request {
	t.Helper()
	b := bytes.NewBufferString(body)
	req := httptest.NewRequest(method, path, b)
	return req
}

func TestContenType(t *testing.T) {
	tests := []struct {
		name                string
		expectedContentType string
		Handler             func(http.Handler) http.Handler
	}{
		{
			name:                "json",
			expectedContentType: "application/json; charset=utf-8",
			Handler:             httpx.ContentTypeJSON,
		},
		{
			name:                "text",
			expectedContentType: "text/plain; charset=utf-8",
			Handler:             httpx.ContentTypeText,
		},
		{
			name:                "html",
			expectedContentType: "text/html; charset=utf-8",
			Handler:             httpx.ContentTypeHTML,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := newRequest(t, http.MethodGet, "/x", "")
			srv := tt.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
			srv.ServeHTTP(rr, req)
			res := rr.Result()
			if value := res.Header.Get("Content-Type"); tt.expectedContentType != value {
				t.Errorf("expected content type %v, got %v", tt.expectedContentType, value)
			}
		})
	}
}
