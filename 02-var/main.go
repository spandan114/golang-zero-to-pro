package main

import (
	"fmt"
	"strconv"
)

func main() {
	var username string = "spandan"
	fmt.Println(username)
	fmt.Printf("Variable type is %T\n", username)

	var isLogedin bool = true
	fmt.Println(isLogedin)
	fmt.Printf("Variable type is %T\n", isLogedin)

	var smallValue uint16 = 256
	fmt.Println(smallValue)
	fmt.Printf("Variable type is %T\n", smallValue)

	var smallFlot float64 = 256.57887766768
	fmt.Println(smallFlot)
	fmt.Printf("Variable type is %T\n", smallFlot)

	//Default value
	var anotherVariable int = 10
	fmt.Println(anotherVariable)
	fmt.Printf("Variable type is %T\n", anotherVariable)

	//Implicit type
	var implicitVar = "Spandan"
	fmt.Println(implicitVar)

	//no var no type style
	novarStyle := "spandan"
	fmt.Println(novarStyle)

	var greeting string = "Hello"
	fmt.Printf("%s %s %d\n", greeting, "spandan", 3000)
	conc, _ := fmt.Printf("%s %s %s\n", greeting, "spandan", strconv.Itoa(3000))
	fmt.Print(conc)
	fmt.Printf("%s\n", greeting)
	fmt.Printf("%x\n", greeting[0]) //ASCII char

}
