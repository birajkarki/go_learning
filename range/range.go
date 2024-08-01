package main

import "fmt"

func main() {

	//range
	// - used to iterate over elements in a variety of data structures
	// - used with arrays, slices, maps, strings, channels
	// - returns index and value
	// - if you only want the value, use _ to ignore the index
	// - if you only want the index, use _ to ignore the value

	// nums := []int{6,7,8,9}
	// sum := 0
	// for _, num := range nums {
	// 	sum += num

	// }
	// println("sum:", sum)

	// nums1 := []int{3,5,6}
	// for _,num := range nums1 {
	// 	println(num)
	// }

	//map range
	// m := map[string]string{"a": "apple", "b": "banana"}
	// for _, v := range m {
	// 	println(v)
	// }

	//unicode code point rune
	// starting byte of rune
	// 255 -> 1 byte, 2 byte

	for i, c := range "golang" {
		// fmt.Println(i, c)
		fmt.Println(i, string(c))
	}

}
