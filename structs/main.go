package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contactInfo: contactInfo{
			email:   "jim.halpert@dundermifflin.com",
			zipCode: 18505,
		},
	}
	jim.updateName("Jimothy")
	jim.print()
}

func (p *person) updateName(newFirstName string)  {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%v", p)
}