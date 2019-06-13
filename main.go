package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "nil")
}

func main() {
	http.HandleFunc("/demand", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
