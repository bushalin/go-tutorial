package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const baseurl string = "http://localhost:8000/"

func main() {
	fmt.Println("web verbs in golang")
	// var geturl string = baseurl + "get"
	// PerformGetRequest(geturl)

	// var postUrl string = baseurl + "post"
	// PerformPostJsonRequest(postUrl)

	var postFormUrl string = baseurl + "postform"
	PerformFormPostRequest(postFormUrl)
}

func PerformGetRequest(param string) {

	response, err := http.Get(param)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("status code: ", response.StatusCode)
	fmt.Println("content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())
	// fmt.Println(string(content))
}

func PerformPostJsonRequest(param string) {
	// fake json payload
	requestBody := strings.NewReader(`
		{
			"country": "japan",
			"coursename": "lets go with golong",
			"price": 0,
			"platform": "learncodeOnline.in"
		}
	`)

	response, err := http.Post(param, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}

func PerformFormPostRequest(param string) {
	//formdata

	data := url.Values{}
	data.Add("firstname", "hasib")
	data.Add("lastname", "hasan")
	data.Add("email", "hasib@go.dev")
	response, err := http.PostForm(param, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	var contentstring strings.Builder
	contentstring.Write(content)

	fmt.Println(contentstring.String())

}
