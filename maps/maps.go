package main

import (
	"fmt"
	"maps"
)

//maps -> hash, dictionary, object

func main() {
	// creating map

	// m := make(map[string]string)

	// // setting an element
	// m["name"] = "Biraj"
	// m["age"] = "22"

	// //getting an element
	// fmt.Println(m["name"], m["age"])

	// //Important : If the key is not present in the map, it will return the zero value of the value type
	// // len function
	// fmt.Println(len(m)) // 2

	// // delete function
	// delete(m, "name")
	// fmt.Println(m) // map[age:22]

	// // clear function
	// clear(m)
	// fmt.Println(m) // map[]



	// m1 := map[string]string{
	// 	"name": "Biraj",
	// 	"age":  "22",
	// 	"city": "Kathmandu",
	// 	"country": "Nepal",
	// }
	// fmt.Println(m1)


	// check if key exists
	// k, ok := m1["name"]
	// fmt.Println(k, ok)
	// if ok {
	// 	fmt.Println("Name exists")
	// } else {
	// 	fmt.Println("Name does not exists")
	// }

m1 := map[string]string{"price" : "1000", "name": "Biraj", "age": "22"}
m2 := map[string]string{"price" : "1000", "name": "Biraj", "age": "22"}

fmt.Println(maps.Equal(m1, m2))

}
