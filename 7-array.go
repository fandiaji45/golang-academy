package main

import (
	"fmt"
)

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars)
	fmt.Println(cars[1])
}
