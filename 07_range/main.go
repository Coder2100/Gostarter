package main

import "fmt"

func main() {
	// create slice
	ids := []int{33, 76, 54, 23, 11, 2}
	//looop through ids

	for i, id := range ids {
		fmt.Printf("%d -ID:  %d\n", i, id)
	}
	//loop not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// range with maps
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sha@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s:  %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
