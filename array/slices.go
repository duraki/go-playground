package main

import "fmt"

/**
Slice is extended interface of array, or key value storage
 */
func main() {

	/** create a slice using make() */
	slice := make([]string, 3)
	fmt.Println("Slice", slice)

	/** assign as with typical arr */
	slice[0] = "You"
	slice[1] = "Can"
	slice[2] = "DO IT!"

	fmt.Println("Slice", slice)
	fmt.Println("Slice count", len(slice))

	/** append to slice */
	slice = append(slice, "bye.")
	fmt.Println("Hello", slice[3])

	/** copy slice to a new one */
	sliceCopy := make([]string, len(slice))
	copy(sliceCopy, slice)
	fmt.Println("Copied slice to sliceCopy", sliceCopy)

	/** slice a slice, heh, syntax `slice[low:high]`, to get dlc between range */
	fmt.Println("Sliced slice", sliceCopy[2:3])

}
