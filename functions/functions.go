package main

import "fmt"

// func add(a , b int) int {
// 	return a + b
// }


// func getLanguage()(string, string, bool) {
// 	return "Go", "Python", true
// }


func processIt() func(a int) int {
	return func(a int)int{
		return 4
	}

}
func main() {
	// fn := func(a int) int {
	// 	return a * 2
	// }
	fn := processIt()
	fmt.Println(fn(2))

	// multiple return values
	// lang1, lang2, _ := getLanguage()
	// fmt.Println(lang1, lang2)


	// result := add(1, 2)
	// fmt.Println(result)
}