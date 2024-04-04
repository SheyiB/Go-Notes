package greetings

import "fmt"

// An app that sent a greeting message for a specified person
func Hello(name string) string {
	// Returns a personalized greeting for the specified person
	message := fmt.Sprintf("Hello %v. Jesus loves you!", name)
	//same as
	//var message string
	//message = fmt.Sprintf("Hello %v. Jesus loves you!", name)
	return message
}
