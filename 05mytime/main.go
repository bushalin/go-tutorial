package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.August, 12, 18, 10, 0, 0, time.UTC)

	fmt.Println(createdDate.Format("02/01/2006 Monday"))
}
