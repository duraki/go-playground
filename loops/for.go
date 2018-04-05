package main

import "fmt"

func main() {
	i := 1

	for i <= 3 { /** normal for loop **/
		fmt.Println("Counter", i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ { /** new var as counter **/
		fmt.Println("Counter", j)
	}

	for {
		fmt.Println("Forever loop ...")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("Counter mod", n)
	}
}
