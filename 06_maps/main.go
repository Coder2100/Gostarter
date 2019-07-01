package main

import "fmt"

func main() {
	// define tdps
	/**
	emails := make(map[string]string)
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sha@gmail.com"
	*/
	// declare and initialize map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sha@gmail.com"}
	fmt.Println(emails)
	// get one value paire
	fmt.Println("Bob's Email: " + emails["Bob"])
	// total of maps
	fmt.Println(len(emails))
	// delete maps
	delete(emails, "Bob")
	fmt.Println(emails)
}
