package main

import (
	"fmt"
	"time"
)

func main() {

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("02-01-2006 15:04 Monday"))

	createdDate := time.Date(2020, time.October, 20, 20, 20, 20, 20, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02-01-2006"))
}
