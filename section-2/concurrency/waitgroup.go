package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Fungsi taskA: butuh 2 detik
func taskA(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Mulai Task A")
	time.Sleep(2 * time.Second)
	fmt.Println("Selesai Task A")
}

// Fungsi taskB: butuh 3 detik
func taskB(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Mulai Task B")
	time.Sleep(3 * time.Second)
	fmt.Println("Selesai Task B")
}

func main() {
	// (Opsional) Batasi jumlah core untuk goroutine
	runtime.GOMAXPROCS(4)

	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(2) // Kita punya 2 task

	go taskA(&wg)
	go taskB(&wg)

	wg.Wait() // Tunggu hingga semua goroutine selesai

	duration := time.Since(start)
	fmt.Println("Total waktu (dengan concurrency):", duration)
}
