package main

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	totalTickets     uint
	remainingTickets uint
	bookings         []UserData
}

var events = make(map[string]*Event)
var wg = sync.WaitGroup{}

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	fmt.Println("Welcome to Tixara: Event Ticket Manager")
	for {
		displayMenu()
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addEvent()
		case 2:
			bookTicket()
		case 3:
			displayBookings()
		case 4:
			deleteEvent()
		case 5:
			fmt.Println("Exiting application. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func displayMenu() {
	fmt.Println("\n--- Main Menu ---")
	fmt.Println("1. Add Event")
	fmt.Println("2. Book Ticket")
	fmt.Println("3. Display All Bookings")
	fmt.Println("4. Delete Event")
	fmt.Println("5. Exit")
	fmt.Print("Choose an option: ")
}

func addEvent() {
	var name string
	var totalTickets uint

	fmt.Print("Enter event name: ")
	fmt.Scan(&name)

	fmt.Print("Enter total tickets: ")
	fmt.Scan(&totalTickets)

	events[name] = &Event{
		totalTickets:     totalTickets,
		remainingTickets: totalTickets,
		bookings:         []UserData{},
	}

	fmt.Printf("Event '%v' added with %v tickets.\n", name, totalTickets)
}

func deleteEvent() {
	if len(events) == 0 {
		fmt.Println("No events available to delete.")
		return
	}

	displayEvents()

	var eventName string
	fmt.Print("Enter the name of the event to delete: ")
	fmt.Scan(&eventName)

	if _, exists := events[eventName]; exists {
		delete(events, eventName)
		fmt.Printf("Event '%v' has been successfully deleted.\n", eventName)
	} else {
		fmt.Println("Invalid event name. Please try again.")
	}
}

func bookTicket() {
	if len(events) == 0 {
		fmt.Println("No events available. Please add an event first.")
		return
	}

	displayEvents()

	var eventName string
	fmt.Print("Enter event name: ")
	fmt.Scan(&eventName)
	event, exists := events[eventName]
	if !exists {
		fmt.Println("Invalid event name. Please try again.")
		return
	}

	fmt.Printf("Remaining tickets for '%v': %v\n", eventName, event.remainingTickets)

	var firstName, lastName, email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	isValidName, isValidEmail, isValidTickets := validateUserInput(firstName, lastName, email, userTickets, event.remainingTickets)

	if isValidName && isValidEmail && isValidTickets {
		event.remainingTickets -= userTickets
		user := UserData{
			firstName:       firstName,
			lastName:        lastName,
			email:           email,
			numberOfTickets: userTickets,
		}
		event.bookings = append(event.bookings, user)

		fmt.Printf("\nThank you, %v %v, for booking %v tickets for '%v'!\n", firstName, lastName, userTickets, eventName)
		fmt.Printf("Total tickets: %v | Remaining tickets: %v\n", event.totalTickets, event.remainingTickets)

		wg.Add(1)
		go sendConfirmation(firstName, lastName, email, userTickets, eventName)
	} else {
		displayValidationError(isValidName, isValidEmail, isValidTickets, event.remainingTickets)
	}
	wg.Wait()
}

func displayEvents() {
	fmt.Println("\nAvailable Events:")
	for name, event := range events {
		fmt.Printf(" - %v: Total tickets: %v, Remaining tickets: %v\n", name, event.totalTickets, event.remainingTickets)
	}
}

func sendConfirmation(firstName, lastName, email string, userTickets uint, eventName string) {
	time.Sleep(10 * time.Second)
	ticketDetails := fmt.Sprintf("%v tickets for %v %v for event '%v'", userTickets, firstName, lastName, eventName)
	fmt.Printf("\n***************\n")
	fmt.Printf("Sending Ticket: %v\nTo: %v\n", ticketDetails, email)
	fmt.Printf("***************\n")
	wg.Done()
}

func displayBookings() {
	if len(events) == 0 {
		fmt.Println("No events available.")
		return
	}

	fmt.Println("\nAll Bookings:")
	for name, event := range events {
		if len(event.bookings) > 0 {
			fmt.Printf("Event: %v\n", name)
			for _, booking := range event.bookings {
				fmt.Printf("  - %v %v (%v tickets)\n", booking.firstName, booking.lastName, booking.numberOfTickets)
			}
			fmt.Printf("Total tickets: %v | Remaining tickets: %v\n", event.totalTickets, event.remainingTickets)
		} else {
			fmt.Printf("Event: %v has no bookings yet.\n", name)
		}
	}
}

func displayValidationError(isValidName, isValidEmail, isValidTickets bool, remainingTickets uint) {
	if !isValidName {
		fmt.Println("Invalid name: First and last names must have at least 2 characters.")
	}
	if !isValidEmail {
		fmt.Println("Invalid email: Must contain '@'.")
	}
	if !isValidTickets {
		fmt.Println("Invalid ticket count: Must be greater than 0 and within available tickets.")
		fmt.Printf("Available tickets: %v\n", remainingTickets)
	}
}
