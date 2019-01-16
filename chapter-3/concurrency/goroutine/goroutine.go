package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func main() {
	// Add a count of two, one for each goroutine
	wg.Add(2)

	fmt.Println("Start Goroutines")
	// Launch a goroutine with label "A"
	go printCounts("A")
	go printCounts("B")
	fmt.Println("Waiting to Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}

func printCounts(label string) {
	// Schedule the call to WaitGroup's Done to tell we are done
	defer wg.Done()
	// Randomly wait
	for count := 1; count <= 10; count++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Count: %d from %s\n", count, label)
	}
}
