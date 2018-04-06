package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time now", time.Now())
	fmt.Println("Unix timestamp", time.Now().Unix())
	fmt.Println("Unix nanosecs", time.Now().UnixNano())
}
