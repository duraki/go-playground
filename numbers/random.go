package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i <= 100; i++ {
		fmt.Println(rand.Intn(100), ",")
	}

	for i := 0; i <= 10; i++ {
		fmt.Println(rand.Float64())
	}
}
