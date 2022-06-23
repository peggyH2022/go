package greetings

import "fmt"

//return a greeting for a named person.
func MyHello(name string) string{
	//return a greeting that embeds the name in a message
	message := fmt.Sprintf("Hello, %v. Nice to meet you!", name)
	return message
}