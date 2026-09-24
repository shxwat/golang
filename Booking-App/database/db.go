package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func InitDB() {
	var err error

	connStr := "host=localhost port=5432 user=shashwat dbname=bookingapp sslmode=disable"

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Error opening connection", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Database is not recheable:", err)
	}
	fmt.Println("Connected to Postgres Successfully")
}
func SavingBooking(firstName string, lastName string, email string, userTickets uint) error {
	insertQuery := `INSERT INTO bookings(first_name, last_name, email, tickets) VALUES($1, $2, $3, $4)`

	_, err := db.Exec(insertQuery, firstName, lastName, email, userTickets)
	return err
}
