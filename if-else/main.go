package main

import (
	"fmt"
)

func main() {
	value := 10
	if (value % 2) == 0 {
		fmt.Println("Angka ini genap wak")
	}

	// cara lain if short statement, short syntax
	// if num := 10; num%2 == 0 { //checks if number is even
	// 	fmt.Println("Angka ini genap wak:", num)
	// 	// return
	// }

	// fmt.Println("Angka ini ganjil")
	// fmt.Println("Angka: ", 292-122)

	// Variabel temporary adalah variabel yang hanya bisa digunakan pada blok seleksi kondisi dimana ia ditempatkan saja
	point := 10000.0
	// Deklarasi variabel temporary hanya bisa dilakukan lewat metode type inference yang menggunakan tanda :=. Penggunaan keyword var disitu tidak diperbolehkan karena akan menyebabkan error.
	calculatePercent := point / 100
	if calculatePercent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", calculatePercent, "%")
	} else if calculatePercent >= 70 {
		fmt.Printf("%.1f%s good\n", calculatePercent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", calculatePercent, "%")
	}
}
