package main

import "fmt"

func main() {
	var arr [10]int /** array of 10 integers **/
	fmt.Println("Total count of array", len(arr))

	arr[5] = 100 /** set value at arr index **/

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Index/Values %s %s\n", i, arr[i])
	}

	moreArr := [5]string{"one", "two", "three", "four", "five"}
	fmt.Println("dlc", moreArr)
}
