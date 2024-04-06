package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	// Set properties of logger which include
	// - Logger Entry prefix and Flag to disable printing
	// - Time, src file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it
	//message, err := greetings.Hello("Elijah")

	//Name Slice
	names := []string{"Elijah", "Oluwaseyifunmi", "Ebenezer", "Mojolajesu", "Mobolaji"}
	//Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	//Use an if statement to check if an error is returned, if yes, print it and exit program
	if err != nil {
		log.Fatal(err)
	}

	// If no error is returned, print the  (message) returned map of messages in the terminal``
	fmt.Println(messages)
}
