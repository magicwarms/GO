package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// kalau mau terganti data nya, pakai cara seperti ini menggunakan tanda pointer *
func changeAddressCountryToIndonesia(address *Address) {
	// func changeAddressCountryToIndonesia(address Address) {
	address.Country = "Indonesia"
	// yang berubah data nya cuman scope di function ini saja
	// fmt.Println(address)
}

func main() {
	// Pointer adalah reference atau alamat memori. Variabel pointer berarti variabel yang berisi alamat memori suatu nilai. Sebagai contoh sebuah variabel bertipe integer memiliki nilai 4, maka yang dimaksud pointer adalah alamat memori dimana nilai 4 disimpan, bukan nilai 4 itu sendiri.
	// Variabel bertipe pointer ditandai dengan adanya tanda asterisk (*) tepat sebelum penulisan tipe data ketika deklarasi.
	// Nilai default variabel pointer adalah nil (kosong). Variabel pointer tidak bisa menampung nilai yang bukan pointer, dan sebaliknya variabel biasa tidak bisa menampung nilai pointer.
	// var number *int
	// var name *string

	// Ada dua hal penting yang perlu diketahui mengenai pointer:
	// - Variabel biasa bisa diambil nilai pointernya, caranya dengan menambahkan tanda ampersand (&) tepat sebelum nama variabel. Metode ini disebut dengan referencing.
	// - Dan sebaliknya, nilai asli variabel pointer juga bisa diambil, dengan cara menambahkan tanda asterisk (*) tepat sebelum nama variabel. Metode ini disebut dengan dereferencing.

	// var numberA int = 4
	// var numberB *int = &numberA

	// fmt.Println("numberA (value)   :", numberA)  // 4
	// fmt.Println("numberA (address) :", &numberA) // 0xc20800a220
	// numberA = 6
	// fmt.Println("numberB (value)   :", *numberB) // 4
	// fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

	// var address1 Address = Address{
	// 	City:     "Batam",
	// 	Province: "Kepulauan Riau",
	// 	Country:  "Indonesia",
	// }
	// ini nama nya pass by value
	// var address2 Address = address1

	// ini nama nya pass by reference, pakai tanda & <-- pointer
	// var address2 *Address = &address1
	// kalau pass by value di golang, yang terganti cuman values address2 saja, address1 tidak terganti
	// kalau pass by reference di golang baru terganti values dari address1 dan address2
	// address2.Country = "England"

	// kalau mau maksain values dari address1 dan address2 sama, pakaikan tanda * <-- pointer di awal variable
	// *address2 = Address{
	// 	City:     "Jakarta",
	// 	Province: "DKI Jakarta",
	// 	Country:  "Indonesia",
	// }
	// fmt.Println(address1)
	// fmt.Println(address2)

	// pointer di function
	var address1 Address = Address{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "England",
	}
	// nah kalo ditambah tanda & <-- pointer baru berubah dia
	changeAddressCountryToIndonesia(&address1)
	fmt.Println(address1)
	// dibawah ini data nya tetep saja tak berubah, karna pass by value
	// changeAddressCountryToIndonesia(address1)
	// fmt.Println(address1)
}
