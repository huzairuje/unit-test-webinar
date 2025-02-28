package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unit-test-webinar/handler"
)

func TestE2E_HelloEndpoint(t *testing.T) {
	// Create a new ServeMux and register the route.
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler.HelloHandler)

	// Start a test server using the mux.
	ts := httptest.NewServer(mux)
	defer ts.Close()

	// Make an HTTP GET request to the /hello endpoint with query parameter "name=Test"
	resp, err := http.Get(ts.URL + "/hello?name=Test")
	if err != nil {
		t.Fatalf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	// Verify that the response has a 200 OK status.
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, resp.StatusCode)
	}

	// Verify that the Content-Type header is "application/json".
	if contentType := resp.Header.Get("Content-Type"); contentType != "application/json" {
		t.Errorf("expected Content-Type 'application/json', got %q", contentType)
	}

	// Decode the JSON response.
	var data map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}

	// Check that the "message" field equals "Hello Test".
	expected := "Hello Test"
	if msg, ok := data["message"]; !ok || msg != expected {
		t.Errorf("expected message %q, got %q", expected, msg)
	}
}
