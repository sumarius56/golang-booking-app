package main

import "fmt"

func main(){

    conferenceName := "GO Conference"
	const conferenceNameTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to, %v  ,booking application\n", conferenceName)
	fmt.Println("There are", remainingTickets, "tickets available")
	fmt.Println("Get your tickets here to attend")



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

	remainingTickets = remainingTickets -  userTickets
	bookings = append(bookings, firstName + " " + lastName )



	fmt.Printf("Hello %v %v, you have requested %v tickets. You will reiceive a confirmation mail at %v!\n", firstName,lastName,userTickets,emailAddress)
	fmt.Println("There are", remainingTickets, "tickets remaining for " , conferenceName, "!")

	fmt.Printf("The participants are : %v\n", bookings)
}


// https://youtu.be/yyUHQIec83I?t=4023
