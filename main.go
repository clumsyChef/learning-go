package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println("Book your tickets for the conference below.")
	// var availableTickets uint8 = 100

	// fmt.Printf("We have %d available.\n", availableTickets)

	// var numberOfTicketToBeBooked uint8 = 0

	// fmt.Print("How many tickets do you want to buy: ")
	// fmt.Scanf("%d", &numberOfTicketToBeBooked)

	// if availableTickets < numberOfTicketToBeBooked {
	// 	fmt.Println("This number of tickets are not available")
	// 	return
	// }

	// var namesOfPeople []string

	// fmt.Println("Please put the names of the people below:")

	// for i := 0; i < int(numberOfTicketToBeBooked); i++ {
	// 	fmt.Printf("%d: ", i+1)
	// 	person := bufio.NewReader(os.Stdin)

	// 	var name string

	// 	name, _ = person.ReadString('\n')

	// 	name = strings.TrimSuffix(name, "\n")

	// 	namesOfPeople = append(namesOfPeople, name)
	// }

	// fmt.Println("Below are the people's whose tickets has been booked:")

	// for ind, name := range namesOfPeople {
	// 	fmt.Printf("%d. %v\n", ind+1, name)
	// }

	// availableTickets -= numberOfTicketToBeBooked

	// fmt.Printf("Currently available tickets: %d\n", availableTickets)

	// infinite loop ticket booking system

	availableTickets := 10

	var listOfPeoples []string

	for {
		var whatToDo string
		var numberOfTickets int

		fmt.Printf("Number of available tickets: %d\n", availableTickets)

		fmt.Println("Select the option: Book tickets (b) or See attendees list (l)")

		fmt.Scan(&whatToDo)

		if whatToDo == "l" {
			for ind, name := range listOfPeoples {
				fmt.Printf("%d. %v\n", ind+1, name)
			}

			continue
		}

		fmt.Print("How many tickets do you want to book: ")

		fmt.Scan(&numberOfTickets)

		if availableTickets < numberOfTickets {
			fmt.Println("Sorry! Can't book more tickets than available.")
			continue
		}

		availableTickets -= numberOfTickets

		fmt.Println("Please provide the names of the attendees below:")

		var currentPeoples []string

		for i := 0; i < numberOfTickets; i++ {
			fmt.Print(i+1, ": ")
			person := bufio.NewReader(os.Stdin)
			name, _ := person.ReadString('\n')
			name = strings.TrimSuffix(name, "\n")
			currentPeoples = append(currentPeoples, name)
		}

		fmt.Println("You have booked tickets for the following people: ")

		for ind, name := range currentPeoples {
			fmt.Printf("%d. %v\n", ind+1, name)
			listOfPeoples = append(listOfPeoples, name)
		}

		fmt.Println("Thank you for the booking, have a nice day :)")

		if availableTickets == 0 {
			fmt.Println("Booking closed")
			fmt.Println("Booking closed")
			fmt.Println("Booking closed")
			fmt.Println("Booking closed")
			return
		}
	}

}
