package main

import "fmt"

// Definisi struct
type Student struct {
	Name  string
	Age   int
	Major string
}

func main() {
	// Inisialisasi struct (cara 1)
	var student1 Student
	student1.Name = "Budi"
	student1.Age = 20
	student1.Major = "Informatika"

	// Inisialisasi struct (cara 2)
	student2 := Student{
		Name:  "Siti",
		Age:   22,
		Major: "Sistem Informasi",
	}

	// Inisialisasi struct (cara 3)
	student3 := Student{"Andi", 21, "Teknik Komputer"}

	// Menampilkan data
	fmt.Println("Mahasiswa 1:", student1)
	fmt.Println("Mahasiswa 2:", student2)
	fmt.Println("Mahasiswa 3:", student3)

	// Akses field
	fmt.Println("Nama Mahasiswa 1:", student1.Name)

	// Update field
	student1.Age = 21
	fmt.Println("Umur terbaru Mahasiswa 1:", student1.Age)

	// Gunakan struct sebagai parameter fungsi
	printStudent(student2)
}

// Fungsi yang menerima struct sebagai parameter
func printStudent(s Student) {
	fmt.Printf("Nama: %s\nUsia: %d\nJurusan: %s\n", s.Name, s.Age, s.Major)
}
