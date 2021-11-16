package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in golang")

	rand.Seed(time.Now().UnixNano())
	rndomNo := rand.Intn(6) + 1
	fmt.Println(rndomNo)

	switch rndomNo {
	case 1:
		fmt.Println("You get 1 , you can open")
	case 2:
		fmt.Println("You get 2 , move 2 spot")
	case 3:
		fmt.Println("You get 3 , move 3 spot")
	case 4:
		fmt.Println("You get 4 , move 4 spot")
		fallthrough //Jump to the next case
	case 5:
		fmt.Println("You get 5 , move 5 spot")
	case 6:
		fmt.Println("You get 6 , move 6 spot")
	default:
		fmt.Println("WTF !")
	}

}
