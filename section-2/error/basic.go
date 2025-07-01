package main

import (
	"errors"
	"fmt"
)

func validasiUmur(umur int) error {
	if umur < 0 {
		return errors.New("UMUR TIDAK BOLEH NEGAIIF")
	}

	if umur > 10 {
		return fmt.Errorf("UMUR  %d BELUM CUKUP UNTUK MENDAFTAR", umur)
	}

	return nil
}

func main() {
	err := validasiUmur(5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Pendaftaran berhasil")
}
