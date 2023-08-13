package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://loc.dev:3000/learn?coursename=reactjs&paymentid=laskdfjla59849"
const testUrl string = "POST https://reqres.in/"

func main() {
	fmt.Println("urls in golang")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The tyep of queryparams are %T\n", qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for key, val := range qparams {
		fmt.Println("param is: ", key, val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitest",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
