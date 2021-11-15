package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	// rand.Seed(time.Now().Unix())
	// fmt.Println(rand.Intn(time.Now().Unix()) + 1)

	nBig, err := rand.Int(rand.Reader, big.NewInt(time.Now().Unix()))
	if err != nil {
		panic(err)
	}

	fmt.Println(nBig)
}
