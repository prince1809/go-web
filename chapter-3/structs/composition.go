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

type Admin struct {
	Person // type embedding for composition
	Roles  []string
}

//override PrintDetails
func (a Admin) PrintDetails() {
	// call person PrintDetails
	a.Person.PrintDetails()
	fmt.Println("Admin Roles:")
	for _, v := range a.Roles {
		fmt.Println(v)
	}
}

type Member struct {
	Person // type embedding for composition
	Skills []string
}

//overrides PrintDetails
func (m Member) PrintDetails() {
	// call person PrintDetails
	m.Person.PrintDetails()
	fmt.Println("Skills:")
	for _, v := range m.Skills {
		fmt.Println(v)
	}
}

func main() {
	alex := Admin{
		Person{
			"Alex",
			"John",
			time.Date(1980, time.January, 10, 0, 0, 0, 0, time.UTC),
			"alex@gmail.com",
			"New York",
		},
		[]string{"Manage Team", "Manage Tasks"},
	}
	prince := Member{
		Person{
			"Prince",
			"Kumar",
			time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			"prince@gmail.com",
			"Tokyo",
		},
		[]string{"Go", "Docker", "kuebernetes"},
	}

	// call the methods for alex
	alex.PrintName()
	alex.PrintDetails()
	// call methods for Prince
	prince.PrintName()
	prince.PrintDetails()
}
