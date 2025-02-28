package main

import (
	"log"
	"net/http"

	"github.com/unit-test-webinar/handler"
)

func main() {
	http.HandleFunc("/hello", handler.HelloHandler)
	log.Println("Starting server on :1234")
	log.Fatal(http.ListenAndServe(":1234", nil))
}
