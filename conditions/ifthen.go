package main

import (
	"fmt"
	"time"
)

func main() {
	if 8-4 == 4 {
		fmt.Println("8-4=4 is", true)
	}

	if bornIn := 1995; time.Now().Year() - bornIn == 23 {
		fmt.Println("You are old enough")
	}
}
