package main

func main() {

	age := 13
	if age >= 18 {
		println("You are an adult!")
	} else if age >= 13 && age < 18 {
		println("You are a minor!")
	} else {
		println("You are a kid!")
	}

}


// go does not have a ternary operator, so if-else is the only way to do conditional logic.


