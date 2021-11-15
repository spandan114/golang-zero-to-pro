package main

import "fmt"

func main() {
	var superhiro [4]string
	superhiro[0] = "Bat man"
	superhiro[1] = "Super man"

	fmt.Println("super hero list is : ", superhiro)
	fmt.Println("super hero 0 is : ", superhiro[0])
	fmt.Println("super hero list length : ", len(superhiro))

	var food = [3]string{"pizza", "burger"}

	fmt.Println(food)
}
