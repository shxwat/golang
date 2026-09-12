package main

import (
	"encoding/json" // Used to convert Go data into JSON
	"fmt"
	"net/http"
)

// STEP 1: CREATE A STRUCT
// This struct will be converted into JSON.
//
// json:"name" is called a JSON tag.
// It decides what key name will appear in the JSON response.

type User struct {
	Name    string `json:"name"`
	Company string `json:"company"`
	Role    string `json:"role"`
}

func main() {

	// STEP 2: CREATE A ROUTE
	// When someone opens "/user", this function will run.

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {

		myProfile := User{
			Name:    "Shashwat",
			Company: "Code Bucket Soltions Pvt Ltd",
			Role:    "SDE 1",
		}

		// STEP 4: SET RESPONSE TYPE
		// Tell the client/browser that the response will be JSON.

		w.Header().Set("Content-Type", "application/json")

		// STEP 5: SEND JSON RESPONSE
		// NewEncoder(w) creates a JSON encoder.
		// Encode(myProfile) converts the struct into JSON
		// and directly sends it to the client.

		json.NewEncoder(w).Encode(myProfile)
	})

	// STEP 6: START THE SERVER

	fmt.Println("Server is running on port 8080....")
	http.ListenAndServe(":8080", nil)
}
