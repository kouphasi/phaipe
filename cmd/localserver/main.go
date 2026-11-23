package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("GET /hello", getHelloHandler)
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/echo/", echoHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func getHelloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	path := r.URL.Path

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "method: %s\n", method)
	fmt.Fprintf(w, "path: %s\n", path)
	fmt.Fprintf(w, "body: %s\n", body)
}
