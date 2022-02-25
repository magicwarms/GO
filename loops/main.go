package main

import "fmt"

func main() {
	// variable yang ada di dalam loop gak bisa dipakai di luar loop
	// for i := 1; i <= 10; i++ {
	// 	if i > 5 {
	// 		break // untuk break looping
	// 	}
	// 	fmt.Println("Nomor berapa nih: ", i)
	// }
	// fmt.Println("Line ini dieksekusi setelah looping")

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 { // odd number bilangan genap
	// 		continue // untuk continue looping
	// 	}
	// 	fmt.Println("Nomor berapa nih: ", i)
	// }
	// fmt.Println("Line ini dieksekusi setelah looping continue")

	// // other way to write loop fx
	// i := 1
	// for i <= 10 {
	// 	fmt.Println("Ke berapa nih:", i)
	// 	i++
	// }

	// outer break
	// outer:
	// 	for i := 0; i < 3; i++ {
	// 		for j := 1; j < 4; j++ {
	// 			// fmt.Print("**")
	// 			fmt.Printf("i = %d , j = %d\n", i, j)
	// 			if i == j {
	// 				break outer
	// 			}
	// 		}
	// 		fmt.Println()
	// 	}

	// for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
	// 	fmt.Printf("%d * %d = %d \n", no, i, no*i)
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// 	if i > 8 {
	// 		break
	// 	}
	// }
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println(counter)
	// 	counter++
	// }

	// di for loop ada short statement
	// seperti dibawah ini
	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Keberapa nih", counter)
	// }

	// cobain iterasi slice
	sliceData := []string{"Andhana", "Mahrani", "Naura", "Elmira"}
	// for i := 0; i < len(sliceData); i++ {
	// 	fmt.Println(sliceData[i])
	// }

	// cobain for range untuk iterasi slice
	for index, value := range sliceData {
		fmt.Println("Index ke:", index, "value nya:", value)
	}
}
