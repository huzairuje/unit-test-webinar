package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

	// Check the Content-Type header.
	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("expected Content-Type 'application/json', got %q", contentType)
	}

	// Decode the JSON response.
	var resp map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("could not unmarshal JSON response: %v", err)
	}

	// Check that the "message" field equals "Hello World".
	expected := "Hello World"
	if msg, ok := resp["message"]; !ok || msg != expected {
		t.Errorf("expected message %q, got %q", expected, msg)
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

	// Check the Content-Type header.
	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("expected Content-Type 'application/json', got %q", contentType)
	}

	// Decode the JSON response.
	var resp map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("could not unmarshal JSON response: %v", err)
	}

	// When no name is provided, the handler defaults to "World".
	expected := "Hello World"
	if msg, ok := resp["message"]; !ok || msg != expected {
		t.Errorf("expected message %q, got %q", expected, msg)
	}
}
