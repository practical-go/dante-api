package main

import "fmt"

func main ()  {
	var conferenceTitle string = "Go Conference"
	const conferenceTickets int = 50 
	var remainingTickets uint = 50 

	fmt.Printf("Welcome to %v  booking application!\n", conferenceTitle);
	fmt.Printf("We have total of %v tickets, and %v are still avaliable\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!");

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email name")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	
}