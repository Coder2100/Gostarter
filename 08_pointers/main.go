package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b) // * represent a pointer in the memory address
	fmt.Println(*&a)
	// change val with pointer
	*b = 10
	fmt.Println(a)
}
