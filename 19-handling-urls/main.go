package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Hello")
	PerformPostFormRequest()
}

func performGetReq() {
	getUrl := "https://jsonplaceholder.typicode.com/todos/1"

	res, err := http.Get(getUrl)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println(res.StatusCode)
	fmt.Println(res.ContentLength)

	var responseString strings.Builder

	result, _ := ioutil.ReadAll(res.Body)
	content, _ := responseString.Write(result)

	fmt.Println("responseString", content)
	fmt.Println(responseString.String())
}

func performPostReq() {
	postUrl := "https://jsonplaceholder.typicode.com/posts"

	reqBody := strings.NewReader(`{
		title: 'foo',
		body: 'bar',
		userId: 1,
	}`)

	res, err := http.Post(postUrl, "application/json", reqBody)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	result, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(result))
}

func PerformPostFormRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/posts"

	//formdata

	data := url.Values{}
	data.Add("title", "hitesh")
	data.Add("body", "choudhary")
	data.Add("userId", "1")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
