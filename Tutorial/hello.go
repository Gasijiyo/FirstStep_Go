package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

// Hello returns a greeing for the name person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
