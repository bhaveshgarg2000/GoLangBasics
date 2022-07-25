package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50

	fmt.Println("Welcome to ", conferenceName)
	fmt.Println("Get your ticktes here")
	fmt.Println("------------------------------------------------------")
	fmt.Println("Conference Name:", conferenceName)
	fmt.Println("Tickets:", conferenceTickets)

	// User Input
	var userName string
	fmt.Print("Enter your name: ")
	// input
	fmt.Scan(&userName)
	// Display on screen
	fmt.Printf("Hello %v", userName)

}
