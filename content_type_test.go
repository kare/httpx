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
	req, err := http.NewRequest(method, path, b)
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func TestContenType(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		Func        func(http.Handler) http.Handler
	}{
		{
			name:        "json",
			contentType: "application/json;charset=utf-8",
			Func:        httpx.ContentTypeJSON,
		},
		{
			name:        "text",
			contentType: "text/plain;charset=utf-8",
			Func:        httpx.ContentTypeText,
		},
		{
			name:        "html",
			contentType: "text/html;charset=utf-8",
			Func:        httpx.ContentTypeHTML,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := newRequest(t, http.MethodGet, "/x", "")
			expectedContentType := tt.contentType
			srv := tt.Func(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
			srv.ServeHTTP(rr, req)
			res := rr.Result()
			if header := res.Header.Get("Content-Type"); expectedContentType != header {
				t.Errorf("expected content type %v, got %v", expectedContentType, header)
			}
		})
	}
}
