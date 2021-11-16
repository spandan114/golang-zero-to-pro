package main

import "fmt"

func main() {

	additionRes := adder(10, 5)
	fmt.Println(additionRes)

	proAdderRes := proAdder(10, 20, 30, 10, 5)
	fmt.Println(proAdderRes)
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}
