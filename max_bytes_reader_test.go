package httpx_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"kkn.fi/httpx"
)

func TestMaxBytesReader(t *testing.T) {
	testData := []struct {
		name     string
		method   string
		maxBytes int
		inputLen int
		err      error
	}{
		{"post request body too large", http.MethodPost, 10, 11, errors.New("http: request body too large")},
		{"put request body too large", http.MethodPut, 10, 11, errors.New("http: request body too large")},
		{"post request body just right", http.MethodPost, 10, 10, nil},
		{"put request body just right", http.MethodPut, 10, 10, nil},
	}
	for _, tc := range testData {
		input := bytes.NewReader(bytes.Repeat([]byte("x"), tc.inputLen))
		req, err := http.NewRequest(tc.method, "/upload", input)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Use(httpx.MaxBytesReader(int64(tc.maxBytes)))
		r.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				if err.Error() != tc.err.Error() {
					t.Fatalf("%v: unexpected error: '%v', expecting '%v'", tc.name, err, tc.err)
				}
			}
			bytesRead := len(body)
			if bytesRead != tc.maxBytes {
				t.Fatalf("%v: expected to read %d bytes, but got %d", tc.name, tc.maxBytes, bytesRead)
			}
		})
		r.ServeHTTP(rr, req)
	}
}
