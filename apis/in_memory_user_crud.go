package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

var userDB = []User{
	{ID: 1, Name: "Shashwat", Role: "SDE 1"},
}

func main() {
	http.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(userDB)
	})

	http.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var newUser User

		json.NewDecoder(r.Body).Decode(&newUser)

		newUser.ID = len(userDB) + 1

		userDB = append(userDB, newUser)

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]any{
			"success":  true,
			"messgage": "New user added",
			"user":     newUser,
		})
	})
	fmt.Println("User API is running on port 9000.")
	http.ListenAndServe(":9000", nil)
}
