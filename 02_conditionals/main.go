package main

import "fmt"

func main() {
	x := 15
	y := 10
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "red"
	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("The prospect have no appeal")
	}
	//switch statement

	switch color {
	case "red":
		fmt.Println("Color is red")
	case "yellow":
		fmt.Println("Color is yellow")
	case "violet":
		fmt.Println("Color is violet")
	default:
		fmt.Println("Im not impressed")
	}
}
