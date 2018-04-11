package main

import (
	"crypto/sha1"
	"fmt"
)

const PASS  =  "password"

func main() {
	s := PASS
	h := sha1.New()

	h.Write([]byte(s))

	sha1 := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("SHA1: %x\n", sha1)
}


