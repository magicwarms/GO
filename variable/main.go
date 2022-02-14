package main

import (
	"fmt"
	"math"
)

func main() {
	// var (
	// 	myAge   int    = 28
	// 	wifeAge string = "TOO OLD"
	// )
	// alternate set variable
	// Perlu diketahui, deklarasi menggunakan := hanya bisa dilakukan di dalam blok fungsi.
	myAge, wifeAge := 28, "TOO OLD"

	fmt.Println("my age is", myAge)
	fmt.Println("my wife age is", wifeAge)
	// short hand declaration
	animal := "Babi"
	fmt.Println(animal, "itu Haram")

	// cobain const di go
	const name = "Lala"
	fmt.Println("Nama aku", name)
	names := "Asmara"
	fmt.Println(names)

	// below code will error
	// Andhana, Rani := 20, 20
	// Tanda := hanya digunakan sekali di awal pada saat deklarasi.
	// Untuk assignment nilai selanjutnya harus menggunakan tanda =
	// Andhana, Rani = 10, 10

	const a, b = 10.45, 12.45
	var calculate = math.Max(a, b)

	fmt.Println("Berapa hayoo", calculate)
	// Underscore (_) adalah reserved variable yang bisa dimanfaatkan untuk menampung nilai
	// yang tidak dipakai.
	// Bisa dibilang variabel ini merupakan keranjang sampah.
	satu, _ := 1, 2
	fmt.Println(satu)

	// di 2022
	// menghitung jumlah string
	var myName string = "Andhana Utama"

	// start count
	var countMyName int = len(myName)
	var getFirstCharacter byte = myName[0]
	fmt.Println()
	fmt.Println(myName)
	fmt.Println(countMyName)
	fmt.Println(getFirstCharacter)
	fmt.Println()
	// kita tidak perlu mendefinisikan tipe data jika langsung define value nya
	// seperti ini
	var hobi = "Coding"
	fmt.Println(hobi)

	var age int = 60
	fmt.Println(age)

	// Motto paling mantap
	var mottoMantap string = "sempardak"
	fmt.Println(mottoMantap + " hehe")
}
