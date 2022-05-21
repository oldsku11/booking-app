package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("\nconferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

}
