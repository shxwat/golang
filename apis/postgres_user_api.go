package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

var db *sql.DB

func main() {
	var err error

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Connection is wrong: ", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Database connection is not building: ", err)
	}

	fmt.Println("PostgreSQL connected successfully.")

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	role TEXT NOT NULL
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Error on creating table: ", err)
	}
	fmt.Println("Users table is ready.")

	http.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`SELECT id, name, role FROM users ORDER BY id`)
		if err != nil {
			log.Println("Error fetching users:", err)
			http.Error(w, "Users could not be fetched", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		users := []User{}
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID, &user.Name, &user.Role); err != nil {
				log.Println("Error reading user:", err)
				http.Error(w, "Users could not be read", http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}

		if err := rows.Err(); err != nil {
			log.Println("Error iterating over users:", err)
			http.Error(w, "Users could not be read", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})

	http.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var newUser User
		json.NewDecoder(r.Body).Decode(&newUser)

		query := `INSERT INTO users (name,role) VALUES ($1, $2) RETURNING id`

		err := db.QueryRow(query, newUser.Name, newUser.Role).Scan(&newUser.ID)

		if err != nil {
			http.Error(w, "User not saved in Database", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"message": "User successfully added to Database",
			"user":    newUser,
		})
	})
	fmt.Println("Server is running on port 9000")
	http.ListenAndServe(":9000", nil)
}
