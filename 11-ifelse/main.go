package main

import "fmt"

func main() {
	fmt.Println("if else in golang")

	myNum := 23

	if myNum > 10 {
		fmt.Println("Greater then 10")
	} else if myNum < 10 {
		fmt.Println("Less then 10")
	} else {
		fmt.Println("Equal to 10")
	}

	if n := 3; n%2 == 0 {
		fmt.Println("enen num")
	} else {
		fmt.Println("odd num")
	}

}
