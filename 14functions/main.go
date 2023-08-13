package main

import "fmt"

func main() {
	fmt.Println("functions in golang")
	greeter()

	result := adder(3, 5)

	proResult := proAdder(3, 5, 6, 3, 2, 5)

	extraAdderResult, err := extraAdder(5, 5)

	fmt.Println("Result is : ", result)
	fmt.Println("Pro Result is : ", proResult)
	fmt.Println("extra adder result is: ", extraAdderResult, err)
}

func greeter() {
	fmt.Println("hello!")
}

// func function_name(parameter1, parameter2...) return_type
func adder(x int, y int) int {
	return x + y
}

// values is now working as a slice
func proAdder(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}
	return total
}
func extraAdder(x int, y int) (int, string) {
	return (x + y), "this is the second string"
}
