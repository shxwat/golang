package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/pokemon/{name}", func(w http.ResponseWriter, r *http.Request) {
		pokemonName := r.PathValue("name")

		fmt.Fprintf(w, "Your favourite pokemon is: %s!", pokemonName)
	})

	fmt.Println(("Server is running on Port 9000"))
	http.ListenAndServe(":9000", nil)
}
