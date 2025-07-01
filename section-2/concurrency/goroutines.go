package main

import (
	"fmt"
	"time"
)

func halo() {
	fmt.Println("Halo dari goroutine!")
}

func main() {
	go halo()                   // Menjalankan fungsi halo() sebagai goroutine
	time.Sleep(1 * time.Second) // Menunggu sebentar agar goroutine sempat dieksekusi
}
