package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// go greetings("Hello")
	// greetings("Spandan")
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go sendReq(web)
		wg.Add(1)
	}
	wg.Wait()
}

// func greetings(s string) {
// 	for i := 0; i <= 6; i++ {
// 		time.Sleep(3 * time.Microsecond)
// 		fmt.Println(s)
// 	}
// }

func sendReq(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d is the response for %s\n", res.StatusCode, endpoint)

}
