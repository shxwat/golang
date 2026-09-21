package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	var bookings [50]string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total off %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get you tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter you email: ")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	// userName = "Tom"

	fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
