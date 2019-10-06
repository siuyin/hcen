package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("ID,Price")
	for i := 1; i <= 500; i++ {
		fmt.Printf("%d,%5.2f\n", i, float32(rand.Int31n(9999)+1)/100.0)
	}
}
