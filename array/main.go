package main

import "fmt"

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)

}

// func printarray(a [3][2]string) {
// 	for _, v1 := range a {
// 		for _, v2 := range v1 {
// 			fmt.Printf("%s ", v2)
// 		}
// 		fmt.Printf("\n")
// 	}
// }

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func main() {
	// var a [3]int //int array with length 3
	// a[0] = 12
	// a[1] = 14
	// a[2] = 16
	// a[3] = 18 <-- bakal error karna yang di deklarasi kan array nya hanya 3
	// fmt.Println(a)
	// fmt.Println("Length nya: ", len(a))

	// another way define array
	// a := [3]int{12, 14, 16}
	// fmt.Println(a[1])
	// It is not necessary that all elements in an array have to be assigned a value during short hand declaration.
	// a := [3]int{12} //kita define length nya 3, tapi kalo kita define 1 aja gak masalah
	// fmt.Println(a)

	// biarkan golang code nya yang cari length nya, pakai ini [...]
	// a := [...]int{12, 14, 16, 17}
	// fmt.Println("data array \t:", a)
	// fmt.Println("jumlah elemen \t:", len(a))
	// a[3] = 18
	// fmt.Println("Ganti data di dalam array 17 menjadi 18 \t:", a)

	// a := [3]int{5, 78, 8}
	// var b [4]int // ini cara deklarasi variable kalo gak ada value nya
	// b = a        // variable a dan b itu adalah tipe yang berbeda, karna size nya beda, dan size array adalah bagian dari tipe di GO
	// fmt.Println(b)

	// Arrays in Go are value types and not reference types

	// a := [...]string{"USA", "China", "India", "Germany", "France"}
	// b := a // a copy of a is assigned to b
	// b[0] = "Singapore"
	// fmt.Println("a is ", a) //original array gak bakal ke ganti tetep sama dia
	// fmt.Println("b is ", b)

	// num := [...]int{5, 6, 7, 8, 8}
	// fmt.Println("before passing to function ", num) //original array
	// changeLocal(num)                                //num is passed by value, walau sudah diganti dengan fn ini
	// fmt.Println("after passing to function ", num)  //original array gak bakal ke ganti tetep sama dia

	// a := []float64{12.22, 17.24, 20.56, 80.44}
	// fmt.Println("Size array nya berapa tuh: ", len(a)) //pakai function len kalo mau tahu size array nya berapa
	// for i := 0; i < len(a); i++ {                      //looping from 0 to the length of the array
	// 	// fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	// 	fmt.Println("Element ke:", i, "Nilainya:", a[i])
	// }

	// sum := float64(0)
	// for i, v := range a { // range returns both the index and value
	// for _, v := range a { // kalo cuman mau value nya aja kayak gitu cara nya kasih underscore (_)
	// fmt.Printf("%d the element of a is %.2f\n", i, v)
	// fmt.Println("Element ke:", i, "Nilainya:", v)
	// 	fmt.Println("Nilainya:", v)
	// 	sum += v
	// 	fmt.Println("Total saat ini:", sum)
	// }
	// fmt.Println("\nsum of all elements of a", sum)

	// array multidimensional
	// a := [3][2]string{
	// 	{"lion", "tiger"},
	// 	{"cat", "dog"},
	// 	{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	// }
	// printarray(a)
	// var b [3][2]string
	// b[0][0] = "apple"
	// b[0][1] = "samsung"
	// b[1][0] = "microsoft"
	// b[1][1] = "google"
	// b[2][0] = "AT&T"
	// b[2][1] = "T-Mobile"
	// fmt.Printf("\n")
	// printarray(b)

	// var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	// // Khusus untuk array yang merupakan sub dimensi atau elemen, boleh tidak dituliskan jumlah datanya. Contohnya bisa dilihat pada deklarasi variabel numbers2
	// var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	// fmt.Println("numbers1", numbers1)
	// fmt.Println("numbers2", numbers2)

	// buat slice pakai keyword make
	// fruits := make([]string, 3)
	// fruits[0] = "apple"
	// fruits[1] = "manggo"
	// fruits[2] = "kelentit"

	// fmt.Println(fruits) // [apple manggo kelentit]
	// fmt.Println("Length nya: ", len(fruits))
	// fmt.Println("Capacity nya: ", cap(fruits))

	// slices in Go Lang
	// ngambil data tertentu di sebuah array
	// kalo code nya gini [] itu berarti slice, kalo array ada isi nya [...] itu array
	// a := [5]int{76, 77, 78, 79, 80}
	// var b []int = a[1:4] //creates a slice from a[1] to a[3]
	// fmt.Println(b)

	// c := []int{6, 7, 8} //creates and array and returns a slice reference
	// fmt.Println(c)

	// months := [...]string{
	// 	"January",
	// 	"February",
	// 	"March",
	// 	"April",
	// 	"May",
	// 	"June",
	// 	"July",
	// 	"August",
	// 	"September",
	// 	"October",
	// 	"November",
	// 	"December",
	// }

	// slice1 := months[4:7]
	// fmt.Println(slice1)
	// fmt.Println("Length nya: ", len(slice1))
	// fmt.Println("Kapasitasnya: ", cap(slice1))
	// // saat rubah slice terubah juga array nya, nih contoh nya
	// // detail penjelasan:
	// // Slices tidak memiliki data sendiri. Ini hanyalah representasi dari array yang mendasarinya. Setiap modifikasi yang dilakukan pada slice akan tercermin dalam array yang mendasarinya.

	// //ubah slice nih dari May ke May Once
	// slice1[0] = "May Once"

	// fmt.Println("ubah slice: ", slice1)
	// // array of months nya juga terubah
	// fmt.Println("Updated months: ", months)

	// starting array 0 and end with n - 1

	// darr := []int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	// fmt.Println("array before", darr)
	// dslice := darr[2:5]
	// for i := range dslice {
	// 	dslice[i]++
	// }
	// fmt.Println("array after", darr)

	// numa := [3]int{78, 79, 80}
	// nums1 := numa[:] //creates a slice which contains all elements of the array
	// nums2 := numa[:]
	// fmt.Println("array before change 1", numa)
	// nums1[0] = 100
	// fmt.Println("array after modification to slice nums1", numa)
	// nums2[1] = 101
	// fmt.Println("array after modification to slice nums2", numa)

	// fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	// fruitslice := fruitarray[1:3]
	// fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6

	// cars := []string{"Ferrari", "Honda", "Ford"}
	// fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	// cars = append(cars, "Toyota")
	// fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	// var names []string //zero value of a slice is nil
	// if names == nil {
	// 	fmt.Println("slice is nil going to append")
	// 	names = append(names, "John", "Sebastian", "Vinay")
	// 	fmt.Println("names contents:", names)
	// }

	// veggies := []string{"potatoes", "tomatoes", "brinjal"}
	// fruits := []string{"oranges", "apples"}
	// food := append(veggies, fruits...) //... <- itu nama nya variadic functions
	// fmt.Println("food:", food)

	// Passing a slice to a function

	// nos := []int{8, 7, 6}
	// fmt.Println("slice before function call", nos)
	// subtactOne(nos)                               // function modifies the slice
	// fmt.Println("slice after function call", nos) // modifications are visible outside

	// learn slices again
	// fruits := []string{"Apel", "Jeruk", "Semangka", "Pisang"}
	// fmt.Println(len(fruits))
	// fmt.Println(cap(fruits))

	// fruitsA := fruits[1:4]
	// fmt.Println(fruitsA)      // [Jeruk Semangka]
	// fmt.Println(len(fruitsA)) // 2
	// fmt.Println(cap(fruitsA)) // 3

	// learn copy slices
	// this source that we want to copy
	source := []string{"watermelon", "pinnaple", "orange", "banana"}
	// and define empty slice
	destination := make([]string, len(source), cap(source))
	copy(destination, source)

	fmt.Println(source)      // watermelon pinnaple
	fmt.Println(destination) // watermelon pinnaple

	// lala := []string{} // deklarasi variabel array
	// fmt.Println(len(lala))

	// colorsArray := [...]string{"green", "red", "yellow"}
	// fruitsSlice := []string{"apple", "grape", "banana", "melon"}
	// newFruits := fruitsSlice[0:3]
	// ini loop array pakai for
	// for i := 0; i < len(colors); i++ {
	// 	item := colors[i]
	// 	println("ini nomor:", i, "value:", item)
	// }

	// ini loop array pakai for-range
	// for i, color := range colors {
	// 	println("ini nomor:", i, "value:", item)
	// }
	// fmt.Println(fruitsSlice)
	// fmt.Println(newFruits)
	// newFruits[2] = "orange"

	// fmt.Println(newFruits)
	// fmt.Println(fruitsSlice)
}
