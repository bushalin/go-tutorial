package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices class")

	// defining a slice
	var fruitList = []string{"Apple", "tomato", "peach"}
	fmt.Println("values of fruitlist: ", fruitList)
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "mango", "banana")
	fmt.Println("values of fruitlist: ", fruitList)

	// this is an example of range
	// range's last value is non-inclusive
	// so the last value is not included in the result
	fruitList = append(fruitList[1:])
	fmt.Println("values of fruitlist: ", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("values of fruitlist: ", fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println("values of fruitlist: ", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 934
	highScores[2] = 434
	highScores[3] = 834
	//highScores[4] = 734

	fmt.Println(highScores)

	highScores = append(highScores, 244, 657, 235)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
