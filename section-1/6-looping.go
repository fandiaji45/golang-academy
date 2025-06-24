package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Iterasi ke:", i)
	}

	fmt.Println("------------------")

	fruits := []string{"Apple", "Banana", "Cherry"}

	for index, value := range fruits {
		fmt.Printf("Index %d: %s\n", index, value)
	}
}
