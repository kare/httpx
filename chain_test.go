//go:build !integration
// +build !integration

package httpx

import (
	"fmt"
	"io"
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
	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error reading response body: %v", err)
	}
	expected := "start\nAPI response\nend\n"
	if string(b) != expected {
		t.Fatalf("expected '%v', but got '%s'", expected, b)
	}
}

func newRequest(method, path string) (*httptest.ResponseRecorder, *http.Request) {
	request, _ := http.NewRequest(method, path, nil)
	recorder := httptest.NewRecorder()

	return recorder, request
}

func BenchmarkChain(b *testing.B) {
	srv := Chain(apiHandler(), headerHandler(), logHandler())
	rr, req := newRequest(http.MethodGet, "/")
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		srv.ServeHTTP(rr, req)
	}
}

func ExampleChain() {
	logHandler := func() Middleware {
		return func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "start\n")
				defer fmt.Fprintf(w, "end\n")
				h.ServeHTTP(w, r)
			})
		}
	}
	headerHandler := func() Middleware {
		return func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("X-Powered-By", "kare")
				h.ServeHTTP(w, r)
			})
		}
	}
	apiHandler := func() http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "API response\n")
		})
	}

	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/index.html", nil)
	srv := Chain(apiHandler(), headerHandler(), logHandler())
	srv.ServeHTTP(rr, r)
	res := rr.Result()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("chain example error: %v\n", err)
	}
	xpb := "X-Powered-By"
	fmt.Printf("%v: %v\n\n", xpb, res.Header.Get(xpb))
	fmt.Print(string(b))
	// Output:
	// X-Powered-By: kare
	//
	// start
	// API response
	// end
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
	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error reading response body: %v", err)
	}
	expected := "start\nAPI response\nend\n"
	if string(b) != expected {
		t.Fatalf("expected '%v', but got '%s'", expected, b)
	}
}
