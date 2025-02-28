package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler_WithName(t *testing.T) {
	// Create a request with a query parameter "name=World"
	req := httptest.NewRequest(http.MethodGet, "/hello?name=World", nil)
	rr := httptest.NewRecorder()

	// Invoke the handler.
	HelloHandler(rr, req)

	// Check for a 200 OK status.
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, rr.Code)
	}

	// Check that the response contains "Hello World".
	expected := "Hello World"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("expected body %q, got %q", expected, rr.Body.String())
	}
}

func TestHelloHandler_WithoutName(t *testing.T) {
	// Create a request without the "name" query parameter.
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rr := httptest.NewRecorder()

	// Invoke the handler.
	HelloHandler(rr, req)

	// Check for a 200 OK status.
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, rr.Code)
	}

	// When no name is provided, the output should be "Hello " (with a trailing space).
	expected := "Hello "
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("expected body %q, got %q", expected, rr.Body.String())
	}
}
