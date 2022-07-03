package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Lets learn about mutex & race condition")

	wg := &sync.WaitGroup{}
	mtx := &sync.RWMutex{}
	score := []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mtx *sync.RWMutex) {
		mtx.Lock()
		fmt.Println("Score 1")
		score = append(score, 1)
		mtx.Unlock()
		wg.Done()
	}(wg, mtx)

	go func(wg *sync.WaitGroup, mtx *sync.RWMutex) {
		mtx.Lock()
		fmt.Println("Score 2")
		score = append(score, 2)
		mtx.Unlock()
		wg.Done()
	}(wg, mtx)

	go func(wg *sync.WaitGroup, mtx *sync.RWMutex) {
		mtx.Lock()
		fmt.Println("Score 3")
		score = append(score, 3)
		mtx.Unlock()
		wg.Done()
	}(wg, mtx)

	go func(wg *sync.WaitGroup, mtx *sync.RWMutex) {
		mtx.Lock()
		fmt.Println(score)
		mtx.Unlock()
		wg.Done()
	}(wg, mtx)

	wg.Wait()
}
