package main

import "strings"

func isValidateUserInput(firstname string,lastname string,email string,userTickets uint) (bool, bool,bool) {
	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName,isValidEmail,isValidTicketNumber
}