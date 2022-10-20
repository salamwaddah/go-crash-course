package main

import (
	"fmt"
	"strconv"
)

// Person define person struct
type Person struct {
	firstName, lastName, gender string
	age                         int
}

// value receiver
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " and I am " + strconv.Itoa(p.age)
}

// pointer receiver
func (p *Person) increaseAge() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "f" {
		p.lastName = spouseLastName
	}
}

func main() {
	// init person using struct
	person1 := Person{
		firstName: "Salam",
		lastName:  "Waddah",
		gender:    "m",
		age:       27,
	}

	// without struct fields
	person2 := Person{
		"Salama",
		"Not Waddah",
		"f",
		27,
	}

	fmt.Println(person1, person2)

	fmt.Println(person1.firstName, person1.age)
	person1.age++
	fmt.Println(person1)

	// using receivers
	person1.increaseAge()
	fmt.Println(person1.greet())

	fmt.Println(person2)
	person2.getMarried(person1.lastName)
	fmt.Println(person2)
}
