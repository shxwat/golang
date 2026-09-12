package main

import (
	"fmt"
	"net/http" // used to create HTTP servers and APIs
)

func main() {

	//STEP1: CREATE A ROUTE
	//When someone opens "/api", this function will run...
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {

		//w = ResponseWriter
		// we use it to send a response back to the client/browser.

		//r = Request
		// It contains information about the incoming request.

		fmt.Fprint(w, "Server is running. This is the first API (29/08/2026).")
	})

	//STEP2: START THE SERVER

	//":8080" means the server will run on port 8080.
	// ListenAndServer keeps the server running and waiting for requests

	http.ListenAndServe(":8080", nil)
}

// http.HandleFunc() → Route create karta hai
// "/api" → API endpoint / path
// w http.ResponseWriter → Response bhejne ke liye
// r *http.Request → Client ki request ki information
// fmt.Fprint(w, ...) → Client ko data return karta hai
// http.ListenAndServe(":8080", nil) → Server ko port 8080 par start karta hai

// Basically flow: Client → /api → handler function → response → Client.
