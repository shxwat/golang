package controllers

import (
	"fmt"
	"net/http"
	"time"

	"booking-app/database"
	"booking-app/helper"

	"github.com/gin-gonic/gin"
)

var conferenceName = "Go Conference"

const ConferenceTicket = 50

var remainingTickets uint = 50

type BookingRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	UserTickets uint   `json:"userTickets"`
}

func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"conference":       conferenceName,
		"totalTickets":     ConferenceTicket,
		"remainingTickets": remainingTickets,
	})
}
func BookTicket(c *gin.Context) {
	var req BookingRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	if remainingTickets == 0 {
		c.JSON(http.StatusConflict, gin.H{"message": "Tickets are sold out. Please come back next year"})
		return
	}
	if req.UserTickets > remainingTickets {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Only %d tickets are available", remainingTickets),
		})
		return
	}

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(req.FirstName, req.LastName, req.Email, req.UserTickets, remainingTickets)

	if !isValidName || !isValidEmail || !isValidTicketNumber {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed. Check your inputs or available tickets"})
		return
	}

	err := database.SavingBooking(req.FirstName, req.LastName, req.Email, req.UserTickets)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save booking to database"})
		return
	}

	remainingTickets = remainingTickets - req.UserTickets

	go sendTicket(req.UserTickets, req.FirstName, req.LastName, req.Email)

	c.JSON(http.StatusOK, gin.H{
		"message":          fmt.Sprintf("Successfully booked %v tickets for %v", req.UserTickets, req.FirstName),
		"remainingTickets": remainingTickets,
	})
}
func SyncTickets() {
	bookedTickets, err := database.GetBookedTicketCount()
	if err != nil {
		fmt.Println("Error fetching booked ticket from DB", err)
		return
	}
	if bookedTickets >= ConferenceTicket {
		remainingTickets = 0
	} else {
		remainingTickets = ConferenceTicket - bookedTickets
	}

	fmt.Printf("System Synced!!.. Total tickets booked: %v. Remaining: %v\n", bookedTickets, remainingTickets)
}
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("\n################")
	fmt.Printf("Sending ticket:\n%v to email address %v\n", ticket, email)
	fmt.Println("################")
}
