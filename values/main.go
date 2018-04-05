package main

import "fmt"

func main() {
	fmt.Println("string" + "concat") /** this I seen in PHP **/
	fmt.Println("1+1 =", 1+1) /** this is normal **/
	fmt.Println(true && false) /** shall be false **/
	fmt.Println(true || false) /** shall be true **/
	fmt.Println(!true) /** shall not be true **/
}