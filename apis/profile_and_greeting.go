package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name    string `json:"name"`
	Company string `json:"company"`
	Role    string `json:"role"`
}

func main() {

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		myProfile := User{
			Name:    "Shashwat",
			Company: "Code Bucket Solution Pvt Ltd",
			Role:    "SDE 1",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(myProfile)
	})

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		frontendName := r.URL.Query().Get("name")

		if frontendName == "" {
			frontendName = "Broo"
		}
		fmt.Fprintf(w, "Hello %s, welcome to the Go Backend!", frontendName)
	})

	fmt.Println("Server is running on port 8000.")
	http.ListenAndServe(":8000", nil)
}
