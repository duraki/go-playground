package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	p(time.Now().Format(time.RFC3339))
	p(time.Now().Format("3:04PM"))
}
