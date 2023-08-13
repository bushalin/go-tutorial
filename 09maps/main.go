package main

import "fmt"

func main() {
	fmt.Println("welcome to maps class")

	// format for map is make(map[key]value)
	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["CS"] = "C Sharp"

	fmt.Println("list of all languages: ", languages)
	fmt.Println("JS is short for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("list of all languages: ", languages)

	// loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}
	for _, value := range languages {
		fmt.Printf("value is %v\n", value)
	}
}
