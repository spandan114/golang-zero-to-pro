package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hii")

	var superHeros = []string{"super man", "bat man", "spider man", "wonder women"}

	fmt.Println(len(superHeros))
	fmt.Println(superHeros[len(superHeros)-1])
	fmt.Printf("Type %T\n", superHeros)

	superHeros = append(superHeros, "Spider man", "abcd man")
	fmt.Println(superHeros)

	fmt.Println(superHeros[1:])         // 1 to end of arr , [1:3] = 1 to 3
	superHeros = append(superHeros[1:]) //delete 1st element of arr
	fmt.Println(superHeros)

	//memory management

	var schoolScore = make([]int, 4)

	schoolScore[0] = 100
	schoolScore[1] = 200
	schoolScore[2] = 400
	schoolScore[3] = 300
	//definitely return err because we define arr size 4
	//schoolScore[4] = 500

	//Allocate new new memory for 2 new val entire memory allocation happen again
	schoolScore = append(schoolScore, 350, 450)
	fmt.Println(schoolScore)

	//sorting
	sort.Ints(schoolScore)
	fmt.Println(schoolScore)
	fmt.Println(sort.IntsAreSorted(schoolScore))

	//remove val bsed on index

	// fmt.Println(superHeros[:2])
	// fmt.Println(superHeros[2+1:])
	supIndex := 2
	superHeros = append(superHeros[:supIndex], superHeros[supIndex+1:]...)
	fmt.Println(superHeros)

	schIndex := 3
	schoolScore = append(schoolScore[:schIndex], schoolScore[schIndex+1:]...)
	fmt.Println(schoolScore)

}
