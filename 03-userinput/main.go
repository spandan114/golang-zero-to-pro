package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "welcome to userinput"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for pizza : ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Your rating is :", input)

	increasedNum, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Number increased", increasedNum+1)
	}
}
