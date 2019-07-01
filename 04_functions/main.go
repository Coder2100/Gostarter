package main

import "fmt"

func greeting(name string) string {
	return "Hallo" + name
}

//func getSum(num1 int, num2 int) int {
func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Zukile"))
	fmt.Println(getSum(6, 8))
}
