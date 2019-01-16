package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	messages <- "Golang"
	messages <- "Gopher"

	// Receive value from buffered channel
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
