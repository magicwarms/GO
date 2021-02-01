package main

import (
	"fmt"
	"math"
	"strings"
)

// parameter dan return type adalah optional di go lang function
// func calculateBill(qty int, price int) int {
// 	total := qty * price

// 	return total
// }

// other way to create function in golang
// kalo type param nya sama gak usah dibikin kayak param function diatas
func calculateBill(qty, price int) int {
	total := qty * price

	return total
}

// example multiple return value result from function
// func rectProps(length, width float64) (float64, float64) {
// 	area := length * width
// 	perimeter := (length + width) * 2
// 	return area, perimeter
// }

// other way to create and function declaration as return values
func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}

func printMessage(message string, arr []string) {
	nameString := strings.Join(arr, " and ")
	fmt.Println(message, nameString)
}

func divideValue(v1, v2 float64) float64 {
	if v1 == 0 {
		fmt.Println("Nilai pembagian tidak boleh NOL!")
		return 0
	}
	res := v1 / v2
	return res
}

func main() {
	// qty, price := 2, 2000
	// fmt.Println("Total pembayaran kamu:", calculateBill(qty, price))

	// multiple return from function
	// area, perimeter := rectProps(10.8, 5.6)
	// fmt.Printf("Area %f Perimeter %f", area, perimeter)
	bagiDulu := divideValue(7, 2)
	fmt.Println(bagiDulu)
	fmt.Println(math.Pi)
	// names := []string{"Andhana", "Rani"}
	// printMessage("halo", names)
	//kalo mau kembalikan hanya 1 return value dari function pakai _ (underscore)
	// area, _ := rectProps(10.8, 5.6) // perimeter variable is discarded
	// fmt.Printf("Area %f ", area)
}
