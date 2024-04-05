package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// An app that sent a greeting message for a specified person
func Hello(name string) (string, error) {
	//Return an error in the case where name was not supplied
	if name == "" {
		return "", errors.New("empty name")
	}


	// Returns a personalized greeting for the specified person If name is supplied
	message := fmt.Sprintf(randomGreeeting(), name)
	//same as
	//var message string
	//message = fmt.Sprintf("Hello %v. Jesus loves you!", name)
	return message, nil
}

// A Function that Returns one of a set of gretting messages.
//This is selected randomly

func randomGreeeting() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Salutations %v! We greet you!",
		"Blessed are you %v! Highly favoured of God",
		"Hello %v. Jesus loves you!",
	}

	//Return a randomly selcted greeting by specifying a random index
	return formats[rand.Intn(len(formats))]
}