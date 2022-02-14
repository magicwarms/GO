package main

import "fmt"

func main() {
	// konversi integer
	// var integerType8 int8 = 100
	// var integerType16 int16 = int16(integerType8)
	// var integerType32 int32 = int32(integerType8)
	// var integerType64 int64 = int64(integerType8)

	// fmt.Println(integerType8)
	// fmt.Println(integerType16)
	// fmt.Println(integerType32)
	// fmt.Println(integerType64)

	// ambil karakter pertama
	var name string = "andhana utama"
	// byte itu alias dari uint 8
	var firstCharacter byte = name[0]
	// konversi lagi ke string
	var convertString string = string(firstCharacter)

	fmt.Println(firstCharacter)
	fmt.Println(convertString)
}
