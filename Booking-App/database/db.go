package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *sql.DB

func InitDB() {
	var err error

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "5432"
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
	)

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
func GetBookedTicketCount() (uint, error) {
	var totalBooked uint
	query := `SELECT COALESCE(SUM(tickets), 0)FROM bookings`
	err := db.QueryRow(query).Scan(&totalBooked)
	return totalBooked, err
}
