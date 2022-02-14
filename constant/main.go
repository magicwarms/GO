package main

import (
	"fmt"
	"math"
)

func main() {
	// a := 10
	// a = 12

	// fmt.Println("Re-assign variable tanpa const jadi nya berapa: ", a)

	// const b = "Gak bisa di ganti karna ini constant"
	// b = "Diganti nih" //ini bakal error gak bisa di assign kalo dia constant type nya
	// fmt.Println(b)

	square := math.Sqrt(4)
	// code dibawah gak bisa ini bakal error karna di Go-Lang gak bisa set variable yang value nya ada fungsi/proses lagi
	// jadi harus ganti ke tipe variabel (var atau :=)
	// const squareConst = math.Floor(20)
	fmt.Println("kalo di square, jadi berapa hayo: ", square)

	//cara bikin constant bertipe adalah begini:
	// const wifeName string = "mahrani mahir"
	// fmt.Printf("tipe nya %T dan value nya adalah %v \n", wifeName, wifeName)

	// var name = "Andhana Utama kayaraya"
	// type iniString string
	// var newName iniString = "Kayaraya"
	// newName = name //kalo ini error, kalo di golang gak bisa mixed type gitu
	// newName = iniString(name) //jadi harus diginiin wak di convert lagi
	// fmt.Printf("ini type nya %T value nya %v\n", newName, newName)
	// fmt.Println("Jadi nama nya adalah: ", newName)

	//gak error dibawah ini karna constant a tidak bertipe apa-apa, jadi bisa di reassign lagi
	// const c = 5
	// var intVar int = c
	// var int32Var int64 = c
	// var float64Var float64 = c
	// var complex64Var complex64 = c
	// fmt.Println("intVar", intVar, "\nint64Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	// const d = 5.9 / 8
	// fmt.Printf("d's type %T value %v\n", d, d)

	// const e = 2.5 // not allowed as constant 2.5 truncated to integer
	// var f = 5
	// g := float64(f) + e
	// fmt.Printf("g's type %T value %v", g, g)

	// multiple constant
	// const (
	// 	firstName  string = "Elmira"
	// 	middleName string = "Sorra"
	// 	lastName   string = "Utama"
	// )

	// fmt.Println()
	// fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)
}
