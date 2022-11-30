package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)


var conferenceName string = "GO Conference"
const conferenceNameTickets = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	emailAddress string
	userTickets uint
}


func main(){



    greetUsers()

	for remainingTickets > 0 {


		firstName,lastName,emailAddress,userTickets :=getUserInput()

		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName,lastName,emailAddress,userTickets,remainingTickets)


		if  isValidName && isValidEmail && isValidTickets {

			booking(firstName,lastName,emailAddress,userTickets)

			go sendTicket() //send ticket in a goroutine like async in js

            printFirstNames()

			var noTicketsRemaining bool = remainingTickets == 0

			if  noTicketsRemaining {
				fmt.Println("Sorry, there are no more tickets available")
				break
			}
		}else if userTickets > remainingTickets {
			fmt.Println("Sorry,", remainingTickets, "ticket(s) remaining")
			fmt.Println("Please enter a valid number of tickets")
		}else
		{

		  if !isValidName {
			fmt.Println("Please enter a valid name.Must be at least 2 characters long")
		  }
		  if !isValidEmail {
			fmt.Println("Please enter a valid email address")
		  }
		  if !isValidTickets {
			fmt.Println("Please enter a valid number of tickets")
		  }


		}

	}

}

	func greetUsers(){
		fmt.Printf("Welcome to %v  booking application\n", conferenceName)
		fmt.Println("There are", remainingTickets, "tickets available")
		fmt.Println("Get your tickets here to attend")
	}

    func printFirstNames(){
		firstNames := []string{}
			for _, booking := range bookings {


				firstNames = append(firstNames, booking.firstName)

			}

			fmt.Printf("The participants are : %v\n", firstNames)
	}



	func getUserInput( ) (string,string,string,uint){
        var firstName string
		var userTickets uint
		var lastName string
		var emailAddress string

		// ask user for name
		fmt.Println("Please enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email address:")
		fmt.Scan(&emailAddress)

		// ask user for number of tickets
		fmt.Println("Please enter the number of tickets you would like to purchase:")
		fmt.Scan(&userTickets)

		return firstName,lastName,emailAddress,userTickets
	}

	func booking( firstName string, lastName string, emailAddress string, userTickets uint) (){
			remainingTickets = remainingTickets -  userTickets

            //create struct for  user
			var userData = UserData {
				firstName: firstName,
				lastName: lastName,
				emailAddress: emailAddress,
				userTickets: userTickets,
			}

			bookings = append(bookings, userData)
			fmt.Printf("List of bookings is %v\n", bookings)


			fmt.Printf("Hello %v %v, you have requested %v tickets. You will reiceive a confirmation mail at %v!\n", firstName,lastName,userTickets,emailAddress)
			fmt.Println("There are", remainingTickets, "tickets remaining for " , conferenceName, "!")


	}


	func sendTicket(){
		time.Sleep(10 * time.Second)
		var ticket = fmt.Sprintf("Sending ticket(s) to %v\n", bookings[0].emailAddress)
		fmt.Println(ticket)
	}
