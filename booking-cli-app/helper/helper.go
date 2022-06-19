package helper

import (
	"strings"
)

func ValidateUserInput(userName string, email string, userTickets int, reaminingTicket int) (bool, bool, bool) {
	isValidName := len(userName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= reaminingTicket
	return isValidName, isValidEmail, isValidTicketNumber
}
