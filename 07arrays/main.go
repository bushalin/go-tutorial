package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Tomato"

	fmt.Println("fruit list is: ", fruitList)
	fmt.Println("length of the fruit list is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("veg list is: ", vegList)
	fmt.Println("veg list is: ", len(vegList))
}
