package httpx

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func logHandler() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "start\n")
			defer fmt.Fprintf(w, "end\n")
			h.ServeHTTP(w, r)
		})
	}
}

func headerHandler() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Powered-By", "kare")
			h.ServeHTTP(w, r)
		})
	}
}

func apiHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API response\n")
	})
}

func TestChain(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/page.html", nil)
	srv := Chain(apiHandler(), headerHandler(), logHandler())
	srv.ServeHTTP(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status %v, but got %v", http.StatusOK, res.StatusCode)
	}
	if powered := res.Header.Get("X-Powered-By"); powered != "kare" {
		t.Fatalf("expected '%v', but got '%v'", "kare", powered)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error reading response body: %v", err)
	}
	expected := "start\nAPI response\nend\n"
	if string(b) != expected {
		t.Fatalf("expected '%v', but got '%s'", expected, b)
	}
}

func logHandlerFunc() MiddlewareFunc {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "start\n")
			defer fmt.Fprintf(w, "end\n")
			h(w, r)
		}
	}
}

func headerHandlerFunc() MiddlewareFunc {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Powered-By", "kare")
			h(w, r)
		}
	}
}

func apiHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API response\n")
	}
}

func TestChainFunc(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/page.html", nil)
	srv := ChainFunc(apiHandlerFunc(), headerHandlerFunc(), logHandlerFunc())
	srv.ServeHTTP(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status %v, but got %v", http.StatusOK, res.StatusCode)
	}
	if powered := res.Header.Get("X-Powered-By"); powered != "kare" {
		t.Fatalf("expected '%v', but got '%v'", "kare", powered)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error reading response body: %v", err)
	}
	expected := "start\nAPI response\nend\n"
	if string(b) != expected {
		t.Fatalf("expected '%v', but got '%s'", expected, b)
	}
}
