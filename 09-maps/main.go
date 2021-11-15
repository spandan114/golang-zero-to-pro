package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]string)

	myMap["JS"] = "Javascript"
	myMap["PY"] = "Python"
	myMap["GO"] = "Golang"

	fmt.Println(myMap)
	fmt.Println(myMap["JS"])

	//delete
	// delete(myMap, "JS")
	// fmt.Println(myMap)

	for key, val := range myMap {
		fmt.Printf("for key %v value is %v\n", key, val)
	}

}
