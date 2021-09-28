package httpx_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"kkn.fi/httpx"
)

func TestHSTS(t *testing.T) {
	type args struct {
		includeSubDomains bool
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "simple",
			args: args{
				includeSubDomains: true,
			},
			expected: "max-age=63072000; includeSubDomains",
		},
		{
			name: "simple no sub domains",
			args: args{
				includeSubDomains: false,
			},
			expected: "max-age=63072000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := newRequest(t, http.MethodGet, "/x", "")
			srv := httpx.HSTS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}), tt.args.includeSubDomains)
			srv.ServeHTTP(rr, req)
			res := rr.Result()
			if header := res.Header.Get("Strict-Transport-Security"); tt.expected != header {
				t.Errorf("expected header %v, got %v", tt.expected, header)
			}
		})
	}
}
