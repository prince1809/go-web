package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	count := make(chan int)

	wg.Add(2)

	fmt.Println("Start Goroutines")
	// Launch a goroutine with label "A"
	go printCounts("A", count)
	//launch a goroutine with label "B"
	go printCounts("B", count)
	fmt.Println("Channel begin")
	count <- 1
	// wait for goroutines to finish
	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("\nTerminating program")
}

func printCounts(label string, count chan int) {
	defer wg.Done()
	for {
		// Receives message from Channel
		val, ok := <-count
		if !ok {
			fmt.Println("Channel was closed")
			return
		}
		fmt.Printf("Count: %d received from %s \n", val, label)
		if val == 10 {
			fmt.Printf("Channel Closed from %s \n", label)
			// Close the channel
			close(count)
			return
		}
		val++
		// send count back to the other goroutine
		count <- val
	}
}
