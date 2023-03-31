package main

import (
	"fmt"
	"log"
	"strconv"
	// "strings"
)

const hotelRooms uint = 50;
var (
	hotelName string = "Adriatique Hotels";
	remainingRooms uint = 50;
	// bookings [] string;
	bookings []map[string]string;
	emails []map[string]string;
)

func main() {
	

	greetUsers(hotelName, hotelRooms, remainingRooms);
	
	fmt.Printf("We have %d rooms can be booked.\n", remainingRooms);
	fmt.Println("Get your room keys here.");


	type UserData struct {
		firsName string;
		lastName string;
		email string;
		numberOfRooms uint;
	}

	
	// this is an infinite loop => for {}
	// infinite loop with conditions => for <cond> {} e.g.: for true {}
	for (remainingRooms > 0) && (len(bookings) < 50) {
		// var bookings [50] string; => size defined, an array
		 // => size undefined, a slice
		var (
			clientName string;
			clientSurname string;
			clientEmail string;
			clientRooms uint;
		)

	
		fmt.Printf("\nEnter your first name: ");
		fmt.Scan(&clientName);

		fmt.Printf("Enter your last name: ");
		fmt.Scan(&clientSurname);

		fmt.Printf("Enter your email address: ");
		fmt.Scan(&clientEmail);

		fmt.Printf("How many rooms you want to book: ");
		fmt.Scan(&clientRooms);
		
		var isNameValid, isEmailValid, isInputValid = validateUserInput(clientName, clientSurname, clientEmail, clientRooms);
		
		/* s
		imple struct
		var userDataStruct = UserData {
			firsName: clientName,
			lastName: clientSurname,
			email: clientEmail,
			numberOfRooms: clientRooms,
		}
		*/



		/*
		if (clientRooms > remainingRooms) {
			fmt.Println("You can't have rooms more than remaining ones.");
			fmt.Printf("Please enter a number lower than %v: ", remainingRooms);
			fmt.Scan(&clientRooms);
		} */

		if (isNameValid && isEmailValid && isInputValid) {

			
			bookRoom(clientName, clientSurname, clientEmail, clientRooms);
			go sendTicket(clientRooms, clientName, clientSurname, clientEmail);
		
			// bookings[0] = clientName + " " + clientSurname; 
			
			
			// create a map for user
			// map[<key type>]<value type>
			var userData = make(map[string] string);
			// instead of this
			bookings = append(bookings, userData);
/*
			bookings = [
				{clientName: "x", clientSurname: "y", ...},
				{clientName: "z", clientSurname: "t", ...},
				.
				.
				.
			]
*/

			userData["clientName"] = clientName;
			userData["clientSurname"] = clientSurname;
			userData["email"] = clientEmail;
			userData["clientRooms"] = strconv.FormatUint(uint64(clientRooms), 10);
			
			/*
			fmt.Printf("The whole array: %v\n", bookings);
			fmt.Printf("The first value: %v\n", bookings[0]);
			fmt.Printf("Array type:%T\n", bookings);
			fmt.Printf("Array length: %d\n", len(bookings));
			*/
		
			// SLICE => A DYNAMIC ABSTRACTION OF AN ARRAY
		
			// fmt.Printf("The whole slice: %v\n", bookings);
			// fmt.Printf("The first value: %v\n", bookings[0]);
			// fmt.Printf("Array type:%T\n", bookings);
			// fmt.Printf("Array length: %d\n", len(bookings));
			
			// for-each loop

			// var names []string = getFirstNames(bookings);
			// fmt.Printf("The first names slice is: %v", names);
			var emails []string = getEmails(bookings);
			fmt.Printf("The emails are: %v\n", emails);
			var areRoomsOut bool = remainingRooms == 0;
			fmt.Printf("List of bookings:\n%v", bookings);
			if (areRoomsOut) {
				fmt.Println("From now on, we're at capacity. Thanks for interest.");
				break;
			}
		} else {
			if (!isNameValid) {
				log.Fatal("ERROR: First or last name you entered is too short");
				//fmt.Println("First or last name you entered is too short");
			}
			if (!isEmailValid) {
				log.Fatal("ERROR: Email address you entered doesn't contain @ sign");
				// fmt.Println("Email address you entered doesn't contain @ sign");
			}
			if (!isInputValid) {
				log.Fatal("ERROR: Incompatible request for rooms");
				// fmt.Println("Incompatible request for rooms");
			}
			continue;
		}
		
		var city string = "Mexico City";
		switch city {
			case "London":
			
			case "Berlin", "Mexico City":

			case "Tokyo", "Peking":

			case "Singapore":

			case "New York":
			
			default:
				fmt.Println("No valid city selected.");
		}
	}
}

func greetUsers(hotelName string, hotelRooms uint, remainingRooms uint) {
	fmt.Printf("Welcome to %v\n", hotelName);
	fmt.Println("You can book rooms for your needs.");
	fmt.Printf("\n%v, your best choice ever.\n", hotelName); // %s works too
//	fmt.Printf("hotelName is %T, hotelRooms is %T, and remainingRooms is %T\n", hotelName, hotelRooms, remainingRooms);
}

// params outside the parenthesis shows the output type
func getFirstNames(bookings []map[string]string) [] string {
	var firstNames [] string;
	for _, booking := range bookings {
		// var names = strings.Fields(booking); // splits the booking with white space and returns a slice
		firstNames = append(firstNames, booking["clientName"]);
	}

	return firstNames;
}

func getEmails(bookings []map[string]string) [] string {
	var emailsList [] string;
	for _, booking := range bookings {
		emailsList = append(emailsList, booking["email"]);
	}
	return emailsList;
}

