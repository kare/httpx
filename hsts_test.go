package httpx_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"kkn.fi/httpx"
)

func TestHSTS(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "simple",
			expected: "max-age=63072000; includeSubDomains",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := newRequest(t, http.MethodGet, "/x", "")
			srv := httpx.HSTS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
			srv.ServeHTTP(rr, req)
			res := rr.Result()
			if header := res.Header.Get("Strict-Transport-Security"); tt.expected != header {
				t.Errorf("expected header %v, got %v", tt.expected, header)
			}
		})
	}
}
