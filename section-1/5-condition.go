package main

import "fmt"

func main() {
	age := 17
	age = 10

	if age < 17 {
		fmt.Println("Belum cukup umur")
	} else if age == 17 {
		fmt.Println("Pas banget 17 tahun")
	} else {
		fmt.Println("Sudah dewasa")
	}

	fmt.Println("------------------")
	// Switch Sederhana
	day := "Senin"

	switch day {
	case "Senin":
		fmt.Println("Hari kerja dimulai")
	case "Sabtu", "Minggu":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Hari biasa")
	}

	fmt.Println("------------------")
	// Switch tanpa kondisi (mirip if berurutan)
	switch {
	case age < 17:
		fmt.Println("Remaja")
	case age >= 17 && age <= 30:
		fmt.Println("Dewasa muda")
	default:
		fmt.Println("Dewasa")
	}
}
