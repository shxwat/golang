package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

var usersDB = []User{
	{ID: 1, Name: "Shashwat", Role: "SDE-1"},
	{ID: 2, Name: "Batman", Role: "Night Shift SDE"},
}

func main() {
	http.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(usersDB)

	})

	http.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var newUser User

		json.NewDecoder(r.Body).Decode(&newUser)

		newUser.ID = len(usersDB) + 1
		usersDB = append(usersDB, newUser)

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"message": "New user added....",
			"user":    newUser,
		})
	})

	http.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")

		userID, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid ID! Put number", http.StatusBadRequest)
			return
		}

		for _, u := range usersDB {
			if u.ID == userID {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(u)

				return
			}
		}
		http.Error(w, "User not found!!", http.StatusNotFound)
	})

	http.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		userID, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		for index, u := range usersDB {
			if u.ID == userID {

				usersDB = append(usersDB[:index], usersDB[index+1:]...)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]any{
					"success": true,
					"message": "User deleted successfully.",
				})
				return
			}
		}
		http.Error(w, "User not found....", http.StatusNotFound)
	})

	fmt.Println("User API is running on port 9000.")
	http.ListenAndServe(":9000", nil)
}
