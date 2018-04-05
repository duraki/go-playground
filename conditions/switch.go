package main

import (
	"fmt"
	"time"
)

func main() {

	i := 5
	fmt.Println("Estimating current index", i, "as")

	switch i { /** does not require break **/
	case 2:
		fmt.Println("Index is type 2")
	case 3:
		fmt.Println("Index is type 3")
	case 4:
		fmt.Println("Index is type 4")
	case 5:
		fmt.Println("Index is type 5")
	}

	fmt.Println("Estimating current date")
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a weekend, enjoy it")
	default:
		fmt.Println("Work work work work")
	}
	
	fmt.Println("Getting variable type")
	getVarType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("It's a bool")
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("It's a string")
		default:
			fmt.Printf("Not defined, can not resolve %T %s\n", t, t)
		}
	}

	getVarType(1)
	getVarType(true)
	getVarType("nešta tamo")
	getVarType(11.11)

}
