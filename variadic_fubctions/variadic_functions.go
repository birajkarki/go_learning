package main

import "fmt"
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total

}
func main () {
// variadic function -> variable number of arguments to a function
// - can be zero or more
// - must be the last parameter
// - can be of any type
// - can be used with fmt.Println, fmt.Printf, fmt.Sprintf
// - can be used with append function
	// fmt.Println(1, 2, 3, 4, 5)
	nums := []int{1, 2, 3, 4, 5}	
	result := sum(nums...)
	fmt.Println(result)
}