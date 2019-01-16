package main

import (
	"fmt"
	"time"
)

type User interface {
	PrintName()
	PrintDetails()
}

type Person struct {
	FirstName, LastName string
	Dob                 time.Time
	Email, Location     string
}

func (p Person) PrintName() {
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}

func (p Person) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s]\n", p.Dob.String(), p.Email, p.Location)
}

func main() {
	alex := Person{
		"Alex",
		"john",
		time.Date(1980, 2, 1, 0, 0, 0, 0, time.UTC),
		"alex@gmail.com",
		"New York",
	}

	users := []User{alex}

	for _, v := range users {
		v.PrintName()
		v.PrintDetails()
	}
}
