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

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Shashwat", Email: "shashwat11muz@gmail.com"},
		{ID: 2, Name: "Rahul", Email: "rahul@example.com"},
	}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Problem is encodind JSON", http.StatusInternalServerError)
	}
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/users", getUsersHandler)

	fmt.Println("Server is running on PORT", port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		fmt.Println("Server failed:", err)
	}
}
