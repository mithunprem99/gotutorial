package main

import "fmt"

func main() {
	conferenceName := "Go conference" //variables in go thus can be used for only var and not for const 
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("conferenceTicket is %T ,remaining %T,conference name %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v and from that %v remaining\n", conferenceTickets, remainingTickets)

	var userName string
	var userTickets int
	//ask the user for name
	userName = "Mithun"
	userTickets = 2

	fmt.Println("Enter your first name: ")
	fmt.Scan(&userName) //used to take the userinput direct from the user
	
	
	
	// fmt.Println(&userName)
	fmt.Printf("User %v,Booked %v this many tickets", userName, userTickets)
}
