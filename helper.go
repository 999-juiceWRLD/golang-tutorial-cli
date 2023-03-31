package main

import (
	"fmt"
	"strings"
	"time"
)

func validateUserInput(firstName string, lastName string, email string, clientRooms uint) (bool, bool, bool) {
	isNameValid := len(firstName) >= 2 && len(lastName) >= 2;
	isEmailValid := strings.Contains(email, "@");
	isInputValid := clientRooms > 0 && clientRooms <= remainingRooms;

	return isNameValid, isEmailValid, isInputValid;
}

func bookRoom(firstName string, lastName string, mail string, clientRooms uint) {
	fmt.Printf("Hello, %s %s. Thank you for booking %d room(s)\n", firstName, lastName, clientRooms);
			fmt.Printf("You will receive a confirmation email at: %s\n", mail);
		
			remainingRooms -= clientRooms;
			fmt.Println(remainingRooms, "rooms left.");
}

func sendTicket(clientRooms uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second);
	var expression string;
	if (clientRooms == uint(1)) {
		expression = "ticket";
	} else {
		expression = "tickets";
	}
	var ticket = fmt.Sprintf("%v %v for %v %v", clientRooms, expression, firstName, lastName);
	fmt.Println("\n\n\n##########################################");
	fmt.Printf("Sending ticket:\n%v\nto email address:\n%v\n", ticket, email);
	fmt.Println("##########################################");
	fmt.Printf("\n\n");
}