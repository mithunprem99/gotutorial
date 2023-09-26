package main

import "fmt"

func main() {
	conferenceName := "Go conference" //variables in go thus can be used for only var and not for const
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("conferenceTicket is %T ,remaining %T,conference name %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v and from that %v remaining\n", conferenceTickets, remainingTickets)

	var firstname string
	var lastname string

	var email string
	var userTickets uint

	//ask the user for name

	userTickets = 2

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstname) //used to take the userinput direct from the user

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastname)
	fmt.Println("Enter your email: ")

	fmt.Scan(&email)

	fmt.Println("Enter your number of tickets: ")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets
	// fmt.Println(&firstname)
	fmt.Printf("Thankyou %v %v for booking %v tickets . You will receive confirmation email at %v\n",firstname,lastname,userTickets,email)
	fmt.Printf("%v remaining tickets in conference %v",remainingTickets,conferenceName)
}
