package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const ConferenceTicket = 50
	var remainingTickets uint = 50

	var bookings []string //empty slice

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v are still available.\n", ConferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	for {

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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {

			bookings = append(bookings, firstName+" "+lastName)
			remainingTickets = remainingTickets - userTickets

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			//list of bookings
			fmt.Printf("These are all our bookings so far: %v\n", bookings)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}

		}else {
			//error handling
			if !isValidName {
				fmt.Println("Error: first name or last name is too short.")
			}
			if !isValidEmail{
				fmt.Println("Error: Email address doesn't contain @ sign.")
			}
			if !isValidTicketNumber{
				fmt.Println("Error: We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
			}
			fmt.Println("Please try again")
		}

	}

}
