package main

import "fmt"

func main() {
	// Definisi Closure adalah sebuah fungsi yang bisa disimpan dalam variabel.
	// Dengan menerapkan konsep tersebut, kita bisa membuat fungsi didalam fungsi,
	// atau bahkan membuat fungsi yang mengembalikan fungsi.

	// Closure merupakan anonymous function atau fungsi tanpa nama.
	// Biasa dimanfaatkan untuk membungkus suatu proses
	// yang hanya dipakai sekali atau dipakai pada blok tertentu saja.

	var getMinMax = func(n ...int) (min, max int) {
		// var min, max int
		for key, value := range n {
			switch {
			case key == 0:
				min, max = value, value
			case value < min:
				min = value
			case value > max:
				max = value
			}
		}
		return
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3, 5}
	var min, max = getMinMax(numbers...)

	fmt.Println("The data are:", numbers)
	fmt.Println("Minimal number:", min)
	fmt.Println("Maximal number:", max)
}
