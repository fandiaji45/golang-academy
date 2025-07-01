package main

import "fmt"

// Account adalah struct dengan field nama (private) dan Username (public)
type Account struct {
	nama     string // private, hanya bisa diakses di dalam package ini
	Username string // public, bisa diakses dari luar package
}

// SetNama menetapkan nilai ke field private 'nama'
func (u *Account) SetNama(n string) {
	u.nama = n
}

// GetNama mengembalikan nilai dari field private 'nama'
func (u Account) GetNama() string {
	return u.nama
}

func main() {
	// Membuat objek Account
	u := Account{Username: "rahmatrdn"}

	// Mengisi field private 'nama' lewat method SetNama
	u.SetNama("Rahmat Ramadhan Putra")

	// Menampilkan hasil
	fmt.Println("Username:", u.Username)      // Bisa langsung akses karena public
	fmt.Println("Nama lengkap:", u.GetNama()) // Akses nama via method (enkapsulasi)

	// fmt.Println(u.nama) //  Akan error jika berada di file/package lain
}
