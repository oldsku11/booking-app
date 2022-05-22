package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint8 = 50
	var bookings []string

	fmt.Printf("\nconferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	var firstName string
	var lastName string
	var email string
	var userTickets uint8

	for {
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %d tickets, so you can't book %d tickets\n", remainingTickets, userTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("\nThank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n", firstName, lastName, userTickets, email)
		fmt.Printf("%d tickets remaining for %s\n\n", remainingTickets, conferenceName)

		var firstNames []string
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	}
}
