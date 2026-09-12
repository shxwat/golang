package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DevProfile struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

func main() {
	http.HandleFunc("/developer/{name}/{role}", func(w http.ResponseWriter, r *http.Request) {

		devName := r.PathValue("name")
		devRole := r.PathValue("role")

		profile := DevProfile{
			Name: devName,
			Role: devRole,
		}
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(profile)
	})
	fmt.Println("Server is running on Port 9000")
	http.ListenAndServe(":9000", nil)

}
