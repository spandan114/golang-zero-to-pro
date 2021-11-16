package main

import "fmt"

func main() {
	fmt.Println("Loop in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	//Type 1
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// Type 2
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	//Type 3
	// for index, day := range days {
	// 	fmt.Printf("Index is %v & day is %v\n", index, day)
	// }

	// Type 4

	rougeValue := 0

myGoto:
	fmt.Println("triggred goto value")

	for rougeValue < 10 {

		if rougeValue == 2 {
			rougeValue++
			goto myGoto
		}

		if rougeValue == 5 {
			// break

			// rougeValue++ //skip 5
			// continue
		}

		fmt.Println(rougeValue)
		rougeValue++

	}

}
