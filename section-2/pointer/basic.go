package main

import (
	"errors"
	"fmt"
)

func main() {
	// Pointer
	tanpaPointer := 200
	denganPointer := &tanpaPointer

	fmt.Println("Nilai asli :", tanpaPointer)
	fmt.Println("Alamat dari pointer:", denganPointer)
	fmt.Println("Nilai dari pointer  :", *denganPointer)

	// Passing by value
	nilaiA := 100
	nilaiB := 200
	fmt.Println("Hasil Penjumlahan :", hitungPenjumlahan(nilaiA, nilaiB))

	// Function with return error
	panjang := 9.0
	lebar := 3.0

	hasil, err := hitungLuasPersegiPanjang(panjang, lebar)
	if err != nil {
		fmt.Println("Terjadi error:", err)
		return
	}

	fmt.Printf("Panjang: %.2f, Lebar: %.2f\n", panjang, lebar)
	fmt.Printf("Luas persegi panjang: %.2f\n", hasil)
}

// Function biasa (by value)
func hitungPenjumlahan(a int, b int) int {
	return a + b
}

// Function dengan return error (by reference / pointer)
func hitungLuasPersegiPanjang(panjang, lebar float64) (float64, error) {
	if panjang <= 0 || lebar <= 0 {
		return 0, errors.New("panjang dan lebar harus lebih dari 0")
	}

	luas := panjang * lebar
	return luas, nil
}
