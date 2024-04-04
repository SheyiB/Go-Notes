package greetings

import (
	"errors"
	"fmt"
)

// An app that sent a greeting message for a specified person
func Hello(name string) (string, error) {
	//Return an error in the case where name was not supplied
	if name == "" {
		return "", errors.New("empty name")
	}


	// Returns a personalized greeting for the specified person If name is supplied
	message := fmt.Sprintf("Hello %v. Jesus loves you!", name)
	//same as
	//var message string
	//message = fmt.Sprintf("Hello %v. Jesus loves you!", name)
	return message, nil
}
