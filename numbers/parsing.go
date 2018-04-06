package main

import (
	"strconv"
	"fmt"
)

const PI = 1.13
const PI_STRING = "1.13"

func main() {
	f, _ := strconv.ParseFloat(PI_STRING, 64)
	fmt.Println("Parsed PI", f)

	i, _ := strconv.ParseInt("13", 0, 64)
	fmt.Println("Parsed INT", i)

	h, _ := strconv.ParseInt("0XDEADBEEF", 0, 64)
	fmt.Println("Parsed HEX", h)

	a, _ := strconv.Atoi("130")
	fmt.Println("Atoi or base of 10 on string result", a)
}
