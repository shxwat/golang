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

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	newUser.ID = 3

	json.NewEncoder(w).Encode(newUser)
}
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Shshwat", Email: "shashwat@example.com"},
		{ID: 2, Name: "Rahul", Email: "rahul@example.com"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("GET /users", getUserHandler)
	http.HandleFunc("POST /users", CreateUserHandler)

	fmt.Println("Server is running on PORT", port)
	http.ListenAndServe(":"+port, nil)
}
