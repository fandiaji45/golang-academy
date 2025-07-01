package main

import (
	"fmt"
	"time"
)

// Simulasi delay API JNE
func getOngkirJNE(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "JNE: Rp11.000"
}

// Simulasi delay API SiCepat
func getOngkirSiCepat(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "SiCepat: Rp10.800"
}

// Simulasi delay dan error pada API J&T
func getOngkirJNT(ch chan string) {
	time.Sleep(1 * time.Second)

	// Simulasi error
	err := true
	if err {
		ch <- "JNT: ERROR saat mengambil ongkir"
		return
	}

	ch <- "JNT: Rp11.000"
}

func main() {
	// Buat channel untuk masing-masing ekspedisi
	chJNE := make(chan string)
	chSiCepat := make(chan string)
	chJNT := make(chan string)

	// Jalankan semuanya secara concurrent
	go getOngkirJNE(chJNE)
	go getOngkirSiCepat(chSiCepat)
	go getOngkirJNT(chJNT)

	// Ambil hasil dari masing-masing channel
	ongkirJNE := <-chJNE
	ongkirJNT := <-chJNT
	ongkirSiCepat := <-chSiCepat

	// Tampilkan hasilnya
	fmt.Println("Ongkir JNE     :", ongkirJNE)
	fmt.Println("Ongkir SiCepat :", ongkirSiCepat)
	fmt.Println("Ongkir J&T     :", ongkirJNT)
}
