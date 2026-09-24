package main

import (
	"booking-app/helper"
	"database/sql"
	"fmt"

	"sync"
	"time"

	"booking-app/database"
	_ "github.com/lib/pq"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var conferenceName = "Go Conference"

const ConferenceTicket = 50

var remainingTickets uint = 50

// var bookings = make([]UserData, 0) //empty slice

var wg = sync.WaitGroup{}

var db *sql.DB

func main() {
	database.InitDB()
	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		wg.Add(1)

		go sendTicket(userTickets, firstName, lastName, email)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("Error: First name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Error: Email address doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Printf("Error: we have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
			}
			fmt.Println("Please try again....")
		}
	}
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v are still available.\n", ConferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}
func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("\nEnter your firstname: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastname: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	err := database.SavingBooking(firstName, lastName, email, userTickets)
	if err != nil {
		fmt.Println("Failed to save bookings!!", err)
		remainingTickets = remainingTickets + userTickets
		return
	}
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("\n################")
	fmt.Printf("Sending ticket:\n%vto email address %v\n", ticket, email)
	fmt.Println("\n################")

	wg.Done()
}
