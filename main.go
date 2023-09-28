package main

import (
	"fmt"
	"sync"
	"time"

	"strings"
)

var conferenceName = "Go conference" //package level variable that could access from outside the scope
const conferenceTickets = 50

var remainingTickets uint = 50
type UserData struct {
	firstname string
	lastname string
	email string
	numberOfTickets uint 
}
var wg = sync.WaitGroup{} 

func main() {
	// conferenceName := "Go conference" //variables in go thus can be used for only var and not for const
	// const conferenceTickets = 50
	// var remainingTickets uint = 50
	fmt.Printf("conferenceTicket is %T ,remaining %T,conference name %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v and from that %v remaining\n", conferenceTickets, remainingTickets)
	greetUsers()
	// for {

		var firstname string
		var lastname string
		var email string
		var userTickets uint
		var booking = make([]UserData , 0)

		//ask the user for name



		userTickets = 2

		fmt.Println("\nEnter your first name: ")
		fmt.Scan(&firstname) //used to take the userinput direct from the user

		fmt.Println("\nEnter your last name: ")
		fmt.Scan(&lastname)
		fmt.Println("Enter your email: ")

		fmt.Scan(&email)

		fmt.Println("Enter your number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstname, lastname, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets

			// create a  map for user
			// we can declare key value as any thing
			var userData = UserData{
				firstname: firstname,
				lastname: lastname,
				email: email,
				numberOfTickets: userTickets,
			}

			// userData["fistname"] = firstname
			// userData["lastname"] = lastname
			// userData["email"] = email
			// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) //this will convert the int to string 10 is a decimal value

			// booking[0] = firstname + " " + lastname
			booking = append(booking, userData)
			wg.Add(1)
			go sendTickets(userTickets,firstname,lastname,email)

			fmt.Printf("The whole slice %v\n", booking)
			fmt.Printf("The firstvalue slice %v\n", booking[0])
			fmt.Printf("slice type%T\n", booking)
			fmt.Printf("slice type%v\n", len(booking))

			// fmt.Println(&firstname)
			fmt.Printf("Thankyou %v %v for booking %v tickets . You will receive confirmation email at %v\n", firstname, lastname, userTickets, email)
			fmt.Printf("%v remaining tickets in conference %v\n", remainingTickets, conferenceName)
			// printFirstName(booking)
			// firstnames := printFirstName(booking)
			// fmt.Printf("The first name of the bookings are : %v\n", firstnames)
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out")
				// break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is invald")
			}
			if !isValidEmail {
				fmt.Println("Email is invald")
			}
			if !isValidTicketNumber {
				fmt.Println("Ticket number is invald")
			}

			fmt.Printf("Your input data is invalid")

		}
		wg.Wait()

	}

// }
// 
func greetUsers() {
	fmt.Printf("Welcome user to the conference %v ", conferenceName)
}

func printFirstName(booking []string) {
	firstnames := []string{}
	for _, booking := range booking {
		var name = strings.Fields(booking)
		firstnames = append(firstnames, name[0])
	} 
}


func sendTickets(userTickets uint,firstname string , lastname string,email string)  {
	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v ",userTickets,firstname,lastname)
	fmt.Printf("Sending tickets %v to email %v\n",tickets,email )
	wg.Done()
}
