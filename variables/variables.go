package main

import "fmt"


func main() {

	var names string = "John"
	fmt.Println(names)

	// infer
	var name = "John"
	fmt.Println(name)

	var isAdult bool = true
	fmt.Println(isAdult)

	var age int = 30
	fmt.Println(age)
	var price float64 = 10.50
	fmt.Println(price)

	//short hand syntax

	firstName := "John"
	fmt.Println(firstName)
}