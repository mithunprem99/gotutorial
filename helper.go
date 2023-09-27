package main

import "strings"

func ValidateUserInput(firstname string,lastname string,email string,userTickets uint,remainingTickets uint) (bool, bool,bool) {
	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName,isValidEmail,isValidTicketNumber
}

// in go export mainly happens when making the name of function to capital letter