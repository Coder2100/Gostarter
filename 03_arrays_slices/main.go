package main

import "fmt"

func main() {
	//fmt.Println("hello world")
	//Array, need to be fixed length

	var fruitlist [2]string

	// asign values
	fruitlist[0] = "apple"
	fruitlist[1] = "banana"

	// declare and simultanous assign
	fruitlistArray := [2]string{"pineaple", "orange"}

	fruitlistSlice := []string{"apricon", "peach", "berry", "dairly"}

	fmt.Println(fruitlist)
	fmt.Println(fruitlistArray)
	fmt.Println(fruitlistSlice)
	//getting length of an array
	fmt.Println(len(fruitlistSlice))
	// get specified array range
	fmt.Println(fruitlistSlice[1:2])

}
