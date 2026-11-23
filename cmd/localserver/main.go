package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("GET /hello", getHelloHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getHelloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello!")
}
