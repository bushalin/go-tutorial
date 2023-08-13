package main

import "fmt"

func main() {

	// defer will work as LIFO (last in first out)
	defer fmt.Println("hello world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("defers in golang")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
