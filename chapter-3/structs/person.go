package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName, LastName string
	Dob                 time.Time
	Email, Location     string
}

// A person method
func (p *Person) PrintName() {
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}

// A person method
func (p *Person) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s]\n", p.Dob.String(), p.Email, p.Location)
}

func (p *Person) ChangeLocation(newLocation string) {
	p.Location = newLocation
}

func main() {
	p := &Person{
		"Prince",
		"Kumar",
		time.Date(19990, time.January, 1, 0, 0, 0, 0, time.UTC),
		"prince@gmail.com",
		"India",
	}
	p.ChangeLocation("Tokyo")
	p.PrintName()
	p.PrintDetails()
}
