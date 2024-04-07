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
	//message := fmt.Sprintf(randomGreeeting(), name)
	//This would make code fail
	message := fmt.Sprintf(randomGreeeting())
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

// Hellos returns a map that assocaites each of the named people with a greeting message
func Hellos(names []string) (map[string]string, error){
	//A map to assocaite names with mesages
	messages := make(map[string]string)
	// Loop through the recieved slice of names, calling the hello function to get a message
	// for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message with the name
		messages[name] = message
	}

	return messages, nil
}