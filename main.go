package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	var conferenceTickets uint = 50
	var remainingTickets uint = 50
	var userName string
	var userTickets uint
	var bookings [50]string
	bookings[0] = "John"
	bookings[1] = "Alice"
	fmt.Println(bookings)

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
	fmt.Println("Enter you name")
	fmt.Scan(&userName)
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
	fmt.Printf("Number of tickets remaining %v", remainingTickets)

}
