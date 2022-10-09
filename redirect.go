package httpx

import (
	"net/http"
	"net/url"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	if r.URL.Scheme != "https" {
		url := url.URL{
			Scheme:      "https",
			Host:        r.Host,
			Path:        r.URL.Path,
			RawQuery:    r.URL.RawQuery,
			RawFragment: r.URL.RawFragment,
		}
		http.Redirect(w, r, url.String(), http.StatusTemporaryRedirect)
	}
}

// Redirect all HTTP traffic to HTTPS.
func Redirect() http.Handler {
	return http.HandlerFunc(redirect)
}
