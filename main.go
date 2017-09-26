// Simple little HTTP listener for testing purposes

package main

import (
	"io"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, Container!\n")
	})
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
