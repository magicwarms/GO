package main

import (
	"log"
)

type human struct {
	head string
	hand string
	foot string
	eyes int
}

type person struct {
	name string
	age  int
}

type student struct {
	grade  int
	person // inisialisasi struct person di struct student
}

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	log.Println("Hello", customer.Name, "My name is ", name)
}

func main() {
	// Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method),
	// yang dibungkus sebagai tipe data baru dengan nama tertentu.
	// var people human
	// people.eyes = 2
	// people.head = "Kepala"
	// people.hand = "Tangan"
	// people.foot = "Kaki"
	// people.foot = "Kika" // overwrite

	// fmt.Println(people)
	// fmt.Println("Mata nya ada:", people.eyes)

	// cara LAIN inisialisasi struct
	// people2 := human{"Kalapa", "Tungun", "Kaki", 2}
	// pakai cara yang dibawah kayak nya lebih enak
	// people3 := human{
	// 	head: "jason",
	// 	hand: "Tingkak",
	// 	foot: "Kakii",
	// 	eyes: 1,
	// }

	// fmt.Println("human 1 :", people.eyes )
	// fmt.Println("human 2 :", people2.head)
	// fmt.Println("human 3 :", people3)

	// var s1 = student{}
	// s1.name = "wick"
	// s1.age = 21
	// s1.grade = 2

	// fmt.Println("name  :", s1.name)
	// fmt.Println("age   :", s1.age)
	// Khusus untuk properti yang bukan properti asli (properti turunan dari struct lain), bisa diakses dengan cara mengakses struct parent-nya terlebih dahulu, contohnya s1.person.age.
	// fmt.Println("age   :", s1.person.age)
	// fmt.Println("grade :", s1.grade)

	var customerData Customer

	customerData.Name = "Andhana"
	customerData.Address = "Batam"
	customerData.Age = 30

	// fmt.Println(customerData)

	// STRUCT METHOD/FUNCTION
	customerData.sayHello("Jin")

}
