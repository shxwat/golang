package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Shashwat", Email: "Shashwat@example.com"},
	{ID: 2, Name: "Rahul", Email: "Rahul@example.com"},
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Println(w, "Invalid Json", http.StatusBadRequest)
		return
	}
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format, send number", http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "No user", http.StatusNotFound)
}
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var updateData User
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	for i, user := range users {
		if user.ID == id {
			users[i].Name = updateData.Name
			users[i].Email = updateData.Email

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}
	http.Error(w, "User Not Found", http.StatusNotFound)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "User Not Found", http.StatusNotFound)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("GET /users", getUserHandler)
	http.HandleFunc("POST /users", createUserHandler)
	http.HandleFunc("GET /users/{id}", getUserByIDHandler)
	http.HandleFunc("PUT /users/{id}", updateUserHandler)
	http.HandleFunc("DELETE /users/{id}", deleteUserHandler)

	fmt.Println("Serve is running on PORT", port)
	http.ListenAndServe(":"+port, nil)
}
