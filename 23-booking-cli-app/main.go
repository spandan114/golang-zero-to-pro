package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

const conferenceTicket = 50

var conferenceName = "Go conference"
var reaminingTicket = 50
var bookings = make([]UserData, 0)

type UserData struct {
	userName    string
	email       string
	userTickets int
}

func main() {

	greetUser()

	for {

		userName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(userName, email, userTickets, reaminingTicket)

		if isValidName && isValidEmail && isValidTicketNumber {
			reaminingTicket = reaminingTicket - userTickets

			// Create a map
			userData := UserData{
				userName:    userName,
				email:       email,
				userTickets: userTickets,
			}

			bookings = append(bookings, userData)

			// fmt.Printf("Print map : %v \n", userData)
			// fmt.Printf("Slice : %v \n", bookings)
			// fmt.Printf("Type : %T \n", bookings)
			// fmt.Printf("Slice length : %v \n", len(bookings))
			// fmt.Printf("Element 0 : %v \n", bookings[0])

			fmt.Printf("Congratulation %v your %v ticket is confirmed !. You will receive a conformation email at %v \n", userName, userTickets, email)
			fmt.Printf("%v ticket is remaining \n", reaminingTicket)

			var firstNames []string = printFirstName()

			fmt.Printf("First names of booking are : %v \n", firstNames)

			go sendTicket(userName, email, userTickets)

		} else {
			if !isValidName {
				fmt.Println("Invalid user name")
			}
			if !isValidEmail {
				fmt.Println("Invalid email")
			}
			if !isValidTicketNumber {
				fmt.Println("Invalid ticket count")
			}
		}
	}
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Println("We have total", conferenceTicket, "tickets and", reaminingTicket, "are still available.")
}

func printFirstName() []string {
	var firstNames = []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.userName)
	}
	return firstNames
}

func getUserInput() (string, string, int) {
	var userName string
	var email string
	var userTickets int

	fmt.Println("Enter your name : ")
	fmt.Scan(&userName)
	fmt.Println("Enter your email address : ")
	fmt.Scan(&email)
	fmt.Println("Enter no of ticker you want to book : ")
	fmt.Scan(&userTickets)
	return userName, email, userTickets
}

func sendTicket(name string, email string, tickets int) {
	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v ticket for %v", tickets, name)

	fmt.Println("***********")
	fmt.Printf("Sending ticket:\n %v \n to email address: %v\n", ticket, email)
	fmt.Println("***********")
}
