package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserProfie struct {
	Name  string `json:"name"`
	Skill string `json:"skill"`
}

func main() {

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		UserSearch := r.URL.Query().Get("q")
		fmt.Fprintf(w, "Search query: %s", UserSearch)
	})

	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		myData := UserProfie{
			Name:  "Shashwat",
			Skill: "React AND Go Backend",
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(myData)
	})

	fmt.Println("Server is running on port 9000.")
	http.ListenAndServe(":9000", nil)
}
