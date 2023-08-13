package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "GET https://lco.dev"
const anotherurl string = "GET https://restcountries.com/v3.1/all"

func main() {
	fmt.Println("handling web request in golang")

	resposne, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("reposne is of type %T\n", resposne)

	defer resposne.Body.Close() // caller's responsibility to close the connection

	dataBytes, err := ioutil.ReadAll(resposne.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataBytes)

	fmt.Println(content)
}
