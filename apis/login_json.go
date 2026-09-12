package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// type LoginRequest struct {
// 	Email string `json:"email"`
// 	Password string `json:"password"`
// }

// func main(){
// 	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
// 		var userLogin LoginRequest

// 		json.NewDecoder(r.Body).Decode(&userLogin)
// 		fmt.Fprintf(w,"Backend received Email: %s and Password: %s", userLogin.Email, userLogin.Password)
// 	})
// 	fmt.Println("Server is running on Port 9000")
// 	http.ListenAndServe(":9000", nil)

// }

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var userLogin LoginRequest
		json.NewDecoder(r.Body).Decode(&userLogin)

		response := LoginResponse{
			Success: true,
			Message: "Login successful, " + userLogin.Email,
		}
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(&response)
	})
	fmt.Println("Server is running on port 9000")
	http.ListenAndServe(":9000", nil)
}
