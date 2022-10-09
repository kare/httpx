package httpx_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"kkn.fi/httpx"
)

func TestRedirectFromHttpToHttps(t *testing.T) {
	testCases := []struct {
		url      string
		location string
	}{
		{
			url:      "http://kkn.fi",
			location: "https://kkn.fi",
		},
		{
			url:      "http://kkn.fi/",
			location: "https://kkn.fi/",
		},
		{
			url:      "http://kkn.fi/pkg/sub/foo",
			location: "https://kkn.fi/pkg/sub/foo",
		},
		{
			url:      "http://kkn.fi/vanity",
			location: "https://kkn.fi/vanity",
		},
		{
			url:      "http://kkn.fi/vanity?go-get=1",
			location: "https://kkn.fi/vanity?go-get=1",
		},
	}
	for _, tt := range testCases {
		tc := tt
		t.Run("", func(t *testing.T) {
			t.Parallel()
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, tc.url, nil)
			srv := httpx.Redirect()
			srv.ServeHTTP(rec, req)
			res := rec.Result()
			if res.StatusCode != http.StatusTemporaryRedirect {
				t.Errorf("expected response status 301, but got %v", res.StatusCode)
			}
			if tc.location != res.Header.Get("Location") {
				t.Errorf("expected response location '%v', but got '%v'", tc.location, res.Header.Get("Location"))
			}
		})
	}
}
