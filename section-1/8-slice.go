package main

import (
	"fmt"
)

func main() {
	mySlice1 := []int{}
	fmt.Printf("myslice1=%v \n", mySlice1)
	fmt.Printf("length=%d \n", len(mySlice1))

	fmt.Println("-----------------------")

	mySlice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Printf("myslice2=%v \n", mySlice2)
	fmt.Printf("length=%d \n", len(mySlice2))

	fmt.Println()

	mySlice3 := []int{1, 2, 3}
	mySlice4 := []int{4, 5, 6}
	mySlice5 := append(mySlice3, mySlice4...)

	fmt.Printf("myslice5=%v \n", mySlice5)
	fmt.Printf("length=%d \n", len(mySlice5))
}
