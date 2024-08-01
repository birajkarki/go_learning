package main

import (
	"fmt"
	"slices"
)

// slices -> dynamic array
// - no fixed size
// - most used construct in go
// - no need to specify size
// - can grow or shrink
// - reference type
// - can be sliced
// - can be copied
// - can be appended
// - can be passed to functions
// - can be used with range keyword
// - can be used with make function
// - can be used with append function
// - can be used with copy function
// - can be used with delete function
// - can be used with len function
// - can be used with cap function

func main() {

	// uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums == nil) // true
	// fmt.Println(len(nums)) // 0

	// var nums = make([]int, 3, 5)
	// capacity -> maximum numbers of elements can fit
	// fmt.Println(cap(nums)) // [0 0 0]

	// nums = append(nums, 1, 2, 3, 4, 5)
	// fmt.Println(nums) // [0 0 0 1 2 3 4 5]

	// numsok := []int{1, 2, 3, 4, 5}
	// fmt.Println(numsok) // [1 2 3 4 5]

	// var oknum = make([]int, 2,5)
	// fmt.Println(oknum) // [0 0 0]
	// // len
	// fmt.Println(len(oknum)) // 2
	// // cap
	// fmt.Println(cap(oknum)) // 5

	// copy function
	// nums := []int{1, 2, 3, 4, 5}
	// nums2 := make([]int, 5)
	// copy(nums2, nums)
	// fmt.Println(nums2) // [1 2 3 4 5]

	// slice operator

	// nums := []int{1, 2, 3, 4, 5}
	// fmt.Println(nums[0:2]) // [1 2]
	// fmt.Println(nums[1:3]) // [2 3]
	// fmt.Println(nums[2:4]) // [3 4]
	// fmt.Println(nums[3:5]) // [4 5]
	// fmt.Println(nums[:2]) // [1 2]
	// fmt.Println(nums[2:]) // [3 4 5]
	// fmt.Println(nums[:]) // [1 2 3 4 5]

	// slice
	// - reference type
	var nums1 = []int{1, 2, 3, 4, 5}
	var nums2 = []int{1, 0, 3, 4, 5}
	fmt.Println(slices.Equal(nums1, nums2)) 

	// 2d slice
	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums) // [[1 2 3] [4 5 6]]
	
}
