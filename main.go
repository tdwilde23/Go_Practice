package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = make([]User, 0)

type User struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber, isValidInput := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidInput {
			bookTicket(firstName, lastName, email, userTickets)

			wg.Add(1)
			go sendTickets(firstName, lastName, email, userTickets)

			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings are %v\n", firstNames)

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Your email is invalid")
			}
			if !isValidTicketNumber {
				fmt.Printf("We only have %v tickets remaining so you cannot book %v tickets\n", remainingTickets, userTickets)
			}
		}

		var noTicketsRemaining = remainingTickets == 0

		if noTicketsRemaining {
			fmt.Println("All tickets are sold out")
			break
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still left\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		var firstName = booking.firstName
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter your number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	var userData = User{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   userTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("User %v booked %v tickets. Confermation has been sent to %v\n", firstName+" "+lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining\n", remainingTickets)
	fmt.Println(bookings)
}

func sendTickets(firstName string, lastName string, email string, userTickets uint) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending %v to email address %v\n", ticket, email)
	fmt.Println("###############")
	wg.Done()
}
