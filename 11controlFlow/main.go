package main

import "fmt"

func main() {
	fmt.Println("control flow in golang")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "high value user"
	} else {
		result = "lutha user"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is not less than 10")
	}

	// if err != nil {

	// }
}
