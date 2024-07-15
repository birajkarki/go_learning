package main

import "fmt"


const firstName string = "Biraj"
func main() {
	const lastName string = "Karki"
	fmt.Println(firstName + " " + lastName)


	const (
		port = 8080
		host = "localhost"

	)
	fmt.Println(port, host)
}