package handler

import (
	"fmt"
	"net/http"
)

// HelloHandler responds with "Hello World" when the /hello URL is requested.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Query().Get("name"))
}
