package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Shashwat", Email: "shashwat@example.com"},
	{ID: 2, Name: "Rahul", Email: "rahul@example.com"},
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	newUser.ID = len(users) + 1
	users = append(users, newUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("GET /users", getUserHandler)
	http.HandleFunc("POST /users", createUserHandler)

	fmt.Println("Server is running on PORT", port)
	http.ListenAndServe(":"+port, nil)
}
