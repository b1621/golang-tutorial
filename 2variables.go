package main 

import "fmt"

func main(){
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Println("We have total of ", conferenceTickets, "tickets and ", remainingTickets,"are still available")
	// fmt.Println("Get your tickets here to attend")

	//  formating output


	fmt.Printf("Welcome to %s booking application.\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n",conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	

}

