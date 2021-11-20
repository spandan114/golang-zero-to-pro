package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Json")
	// CreateJson()
	ConsumeJson()
}

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"` //omitempty means if value is nil then dont through it
}

func CreateJson() {
	myData := []course{
		{"Mern bootcamp", 499, "Web.com", "abcd", []string{"web-dev", "full-stack"}},
		{"Vue bootcamp", 299, "Web.com", "abcd", []string{"web-dev"}},
		{"Angular bootcamp", 299, "Web.com", "abcd", nil},
	}
	convert, err := json.MarshalIndent(myData, "", "\t") // "/t" = tab
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", convert)

}

func ConsumeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "web.in",
		"tags": ["web-dev","js"]
	}
	`)

	//validate & extract in a struct form
	var myCourse course

	checkValidity := json.Valid(jsonDataFromWeb)

	if checkValidity {
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("JSON data is not valid")
	}

	//Extract key val pair

	var myJsonData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myJsonData)
	fmt.Printf("%#v\n", myJsonData)

	for k, v := range myJsonData {
		fmt.Printf("key %v & value %v\n", k, v)
	}

}
