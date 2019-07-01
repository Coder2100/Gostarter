package main

import (
	"fmt"
	"strconv"
)

//define person struck
type Person struct {
	firstName string
	lastName  string
	age       int
	gender    string
	// shorthand
	/**
	firstName,lastName,gender string
	age int*/
}

//methods go outside of the class

//Greeting method (value receiver)

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// Value changer/ pointer receiver uses * istrikcs
func (p *Person) hasBirthBirthday() {
	p.age++
}

// pointer receiver
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	//initialize person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", age: 36, gender: "f"}

	//alternative
	person2 := Person{"Samantha", "Smith", 36, "f"}
	person3 := Person{"Odwa", "Sammy", 20, "m"}
	fmt.Println(person1)
	fmt.Println(person2)
	// acess keyvalue pairs
	fmt.Println(person1.firstName)
	/**change the struck
	person1.age++
	**/

	//apply get married
	person1.getMarried("Mtotso")
	// has birthday point receiver

	person1.hasBirthBirthday()
	person1.hasBirthBirthday()
	person1.hasBirthBirthday()
	fmt.Println(person1.age)
	fmt.Println(person1.greet())
	fmt.Println(person3.greet())

	//fmt.Println(person1.hasBirthBirthday())

}
