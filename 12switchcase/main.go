package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch case in golang")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is one and you can open")
	case 2:
		fmt.Println("Dice value is one and you can move 2 spots")
	case 3:
		fmt.Println("Dice value is one and you can move 3 spots")
		fallthrough
	case 4:
		fmt.Println("Dice value is one and you can move 3 spots")
	case 5:
		fmt.Println("Dice value is one and you can move 3 spots")
	case 6:
		fmt.Println("you can move 6 spots and you can roll again")
		fallthrough
	default:
		fmt.Println("what the hell!?")
	}
}
