package main

import "fmt"

func main() {
	var conferenceName = "Go conference" //variables in go
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to our %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v and from that %v remaining\n",conferenceTickets,remainingTickets)

}
