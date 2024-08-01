package main

import "fmt"

// closures -> anonymous functions that can access variables defined in the parent function


func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
func main (){
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())



}