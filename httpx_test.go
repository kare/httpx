package httpx

import (
	"context"
	"net/http"
	"testing"
)

type handler struct {
	HandlerWithContext
	isCalled bool
}

func (h *handler) ServeHTTPWithContext(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	h.isCalled = true
	return nil
}

func TestHandlerWithContextAPI(t *testing.T) {
	h := handler{}
	if err := h.ServeHTTPWithContext(context.TODO(), nil, nil); err != nil {
		t.Errorf("error on ServeHTTPWithContext(): %v", err)
	}
	if !h.isCalled {
		t.Error("HandlerWithContext.ServeHTTPWithContext() was not called")
	}
}
