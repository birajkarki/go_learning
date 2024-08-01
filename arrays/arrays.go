package main

import "fmt"

func main() {


	// int -> 0 , string -> "", bool -> false
	//numbered sequence of specific length
	// var nums [4]int

	//assign values to array
	// nums[0] = 1
	// nums[1] = 2
	// nums[2] = 3
	// nums[3] = 4

	//print array
	// fmt.Println(nums) // [1 2 3 4]

	//array length
	// fmt.Println(len(nums)) // 4

	//declare and assign values to array
	// nums := [4]int{1, 2, 3, 4}
	// fmt.Println(nums) // [1 2 3 4]

	// var fruits [3]string
	// fruits[0] = "apple"
	//  fmt.Println(fruits) // apple


	//  var bol [3]bool
	//  bol[0] = true
	//  fmt.Println(bol) // [true false false]


	// 2d array
	nums := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(nums) // [[1 2] [3 4]]

}


// - fixed size, that is predictable
// - memory optimization
// - contant time access