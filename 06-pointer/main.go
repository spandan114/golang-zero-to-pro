package main

import "fmt"

func main() {
	fmt.Println("hii")
	// var ptr *int = 10
	// fmt.Println(ptr)

	myval := 10
	var ptr = &myval
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr + 6
	fmt.Println(myval)

}
