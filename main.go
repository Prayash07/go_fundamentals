package main

import "fmt"

func main() {

	// Datatypes Initialization
	var conferenceName = "Go Conference"
	var conferenceTickets int = 50
	remainingTickets := 50
	var userName string
	var userTickets int

	// array
	var bookings [50]string
	bookings[0] = "John"
	bookings[1] = "Alice"
	fmt.Println(bookings)

	// slice
	bookingSlice := []string{}
	bookingSlice = append(bookingSlice, "John Doe", "Aliice")
	fmt.Println(bookingSlice)

	// loops
	for index, booking := range bookingSlice {
		fmt.Printf("index %v, name %v \n", index, booking)
	}

	for _, booking := range bookingSlice {
		fmt.Println(booking)
	}

	// input and output
	// show data
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	// to get data from user
	fmt.Println("Enter you name")
	fmt.Scan(&userName)
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	if userTickets > remainingTickets {
		fmt.Printf("Tickets not available. We have %v tickets remaining", remainingTickets)
	}

	// logic
	remainingTickets = remainingTickets - userTickets

	// show result

	// if statement
	if remainingTickets <= 0 {
		fmt.Println("Tickets is not available")
	} else if remainingTickets > 0 {
		fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
		fmt.Printf("Number of tickets remaining %v", remainingTickets)
	} else {
		fmt.Println("Else statement")
	}

	// switch statements
	cityName := "Kathmandu"
	switch cityName {
	case "Kathmandu":
		fmt.Println("Kathmandu Selected")
	case "Bhaktapur":
		fmt.Println("Bhaktapur Selected")
	default:
		fmt.Println("Not given")

	}

}
