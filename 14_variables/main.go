package main

import "fmt"

func main() {
	//fmt.Println("Hello world")
	//MAIN TYPES
	// USING var
	var name string = "Zukile"
	var age int = 35
	// shorthand
	//name := "Zukile"
	venue, email := "Cape Town", "zmtotso@gmail.com"
	fmt.Println(name, age, venue, email)
	fmt.Printf("%T\n", name)
}
