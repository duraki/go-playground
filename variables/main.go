package main

import (
	"fmt"
	"time"
)

const BORN_IN = 1995
const IS_SMOKER = true

func main() {
	var fullName = "Halis Duraki" /** 1st way to assign val to @var: `var x = "data"` **/
	fmt.Println("Your name:", fullName)

	var age = time.Now().Year() - BORN_IN
	fmt.Println("Your age:", age)

	fmt.Println("Do you smoke:", IS_SMOKER)


	/** 2nd way to assign vtv is shorthand pascal-like version (ah, good old delphi) */
	newVariableString := "thanks for using"
	fmt.Println("This is assigned and set to correct type,", newVariableString)

}
