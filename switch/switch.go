package main
// import "time"
func main() {

	// age := 13
	// switch {
	// case age >= 18:
	// 	println("You are an adult!")
	// case age >= 13 && age < 18:
	// 	println("You are a minor!")
	// default:
	// 	println("You are a kid!")
	// }

	// switch can also be used to compare values:
	// switch age {
	// case 13:
	// 	println("You are 13 years old!")
	// case 14:
	// 	println("You are 14 years old!")
	// case 15:
	// 	println("You are 15 years old!")
	// default:
	// 	println("You are not 13, 14, or 15 years old!")
	// }



	// time := time.Now()
	// switch {
	// case time.Hour() < 12:
	// 	println("Good morning!")
	// case time.Hour() < 17:
	// 	println("Good afternoon!")
	// default:
	// 	println("Good evening!")
	// }

	// type switch
	var x interface{}
	x = 10
	switch x.(type) {
	case int:
		println("x is an int")
	case string:
		println("x is a string")
	default:
		println("x is another type")
	}

	
}