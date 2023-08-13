package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("handling json in golang")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"React", 299, "learncodeonline.in", "abc123", []string{"web-dev", "dev"}},
		{"angular", 199, "learncodeonline.in", "abc123", []string{"full-stack-dev", "dev", "js"}},
		{"vue", 399, "learncodeonline.in", "abc123", nil},
	}

	//package this data as JSON data
	// finalJson, err := json.Marshal(lcoCourses)

	// \t is for indenting as json format
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
        {
            "coursename": "ReactJS Bootcamp",
            "price": 299,
            "website": "learncodeonline.in",
            "tags": [
                "web-dev",
                "js"
            ]
        }
    `)

	var lcoCourse course
	checkvalid := json.Valid(jsonDataFromWeb)

	if checkvalid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	//some cases where you just want to add data to key value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}
