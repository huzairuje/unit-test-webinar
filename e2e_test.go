package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
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

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	// Check that the response is "Hello Test".
	expected := "Hello Test"
	if strings.TrimSpace(string(body)) != expected {
		t.Errorf("expected body %q, got %q", expected, string(body))
	}
}
