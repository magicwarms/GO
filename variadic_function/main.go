package main

import (
	"fmt"
)

func calculate(numbers ...int) (int, float64) {
	total := 0
	for _, number := range numbers {
		total += number
	}
	average := float64(total) / float64(len(numbers))
	return total, average
}

// variadic function mesti ditempatkan paling akhir posisi nya
// func myHobbies(name string, hobbies ...string) {
// 	hobbiesJoin := strings.Join(hobbies, ", ")

// 	fmt.Println("Nama ku adalah:", name)
// 	fmt.Println("Hobbi ku merupakan:", hobbiesJoin)
// }

func main() {
	// Go mengadopsi konsep variadic function atau pembuatan fungsi dengan parameter sejenis yang tak terbatas. Maksud tak terbatas disini adalah jumlah parameter yang disisipkan ketika pemanggilan fungsi bisa berapa saja.
	// Parameter variadic memiliki sifat yang mirip dengan slice. Nilai dari parameter-parameter yang disisipkan bertipe data sama, dan ditampung oleh sebuah variabel saja.

	// numbers := []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	// totalNumber, averageNumber := calculate(numbers...)
	// cara atas sama bawah sama aja
	totalNumber, averageNumber := calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	fmt.Println(averageNumber, totalNumber)

	// myHobbies("Andhana", "Coding", "Makan")
	// Jika parameter kedua dan seterusnya ingin diisi dengan data dari slice, maka gunakan tanda titik tiga kali.
	// var hobbies = []string{"sleeping", "eating"}
	// myHobbies("wick", hobbies...)
	// return
}
