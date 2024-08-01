package main

func main() {

	//range
	// - used to iterate over elements in a variety of data structures
	// - used with arrays, slices, maps, strings, channels
	// - returns index and value
	// - if you only want the value, use _ to ignore the index
	// - if you only want the index, use _ to ignore the value
	


	nums := []int{6,7,8,9}
	sum := 0
	for _, num := range nums {
		sum += num

	}
	println("sum:", sum)
}