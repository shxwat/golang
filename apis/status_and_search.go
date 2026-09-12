package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Server is running.")
	})
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {

		userSearch := r.URL.Query().Get("q")

		fmt.Fprintf(w, "Search query: %s", userSearch)
	})

	fmt.Println("Server is running on port 9000")

	http.ListenAndServe(":9000", nil)
}
