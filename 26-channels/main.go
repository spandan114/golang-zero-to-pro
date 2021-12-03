package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hii")

	myChan := make(chan int)
	wg := &sync.WaitGroup{}
	// myChan <- 8
	// fmt.Println(<-myChan)
	wg.Add(2)

	// receive ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, isChanelOpen := <-myChan

		fmt.Println(isChanelOpen)
		fmt.Println(val)

		// fmt.Println(<-myChan)

		wg.Done()
	}(myChan, wg)

	// send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChan <- 10
		close(myChan)
		// myChan <- 6
		wg.Done()
	}(myChan, wg)

	wg.Wait()

}
