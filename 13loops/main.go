package main

import "fmt"

func main() {
	fmt.Println("loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// other language style
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// range style
	// i will return index
	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("index is %v and the value is %v\n", index, day)
	}

	rogueValue := 1

	for rogueValue < 10 {
		if rogueValue == 2 {
			goto lco
		}
		if rogueValue == 5 {
			break
		}
		fmt.Println("value is: ", rogueValue)
		rogueValue++
	}

	// goto statement
lco:
	fmt.Println("jumping at learcodeonline.in")
}
