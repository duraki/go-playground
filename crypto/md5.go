package main

import (
	"crypto/md5"
	"fmt"
)

const PASS = "password"

func main() {
	s := PASS
	h := md5.New()

	h.Write([]byte(s))
	md5 := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("MD5: %x\n", md5)
}
