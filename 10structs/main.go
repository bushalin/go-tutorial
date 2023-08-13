package main

import "fmt"

func main() {
	fmt.Println("structs are alternatives to class. Welcome")

	// there is no inheritance in golang; No super or parent

	hasib := User{"Hasib", "hasib@go.dev", true, 16}
	fmt.Println(hasib)
	fmt.Printf("hasib details are: %+v\n", hasib)
	fmt.Printf("Name is %v and Email is %v\n", hasib.Name, hasib.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
