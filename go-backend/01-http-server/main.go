package main

import (
	"fmt"
	"net/http"
	"os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to Go Backend, Shashwat!!")
}
func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", homeHandler)
	fmt.Println("Server is running on PORT", port)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		fmt.Println("Server Crashed:", err)
	}

}
