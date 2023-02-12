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
	bookingSlice = append(bookingSlice, "Prayash")
	fmt.Println(bookingSlice)

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

	// logic
	remainingTickets = remainingTickets - userTickets

	// show result
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
	fmt.Printf("Number of tickets remaining %v", remainingTickets)

}
