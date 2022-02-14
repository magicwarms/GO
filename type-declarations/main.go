package main

import "fmt"

func main() {
	// type declarations biasa nya digunakan untuk membuat alias terhadap tipe data yang sudah ada,
	// dengan tujuan agar lebih dimengerti saja
	// contoh:

	type noKTP string
	type isMarried bool

	var nomorKTP noKTP = "1111111"
	var married isMarried = true

	fmt.Println(nomorKTP)
	fmt.Println(married)

	// untuk overwriting value
	fmt.Println(noKTP("12345678910"))

}
