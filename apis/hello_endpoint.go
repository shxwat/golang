package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "System is ready.")
	})
	fmt.Println("Server is running on 8080....")

	http.ListenAndServe(":8080", nil)
}
