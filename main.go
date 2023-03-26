package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/docker", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello from Docker container")
	})

	http.ListenAndServe(":8080", nil)
}
