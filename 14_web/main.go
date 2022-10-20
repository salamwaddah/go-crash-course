package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go!")
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}
