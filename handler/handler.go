package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HelloHandler responds with a JSON object containing a greeting message.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	response := map[string]interface{}{
		"message": fmt.Sprintf("Hello %s", name),
	}

	// Set the header to indicate JSON response.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Marshal the response map into JSON and write it.
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
