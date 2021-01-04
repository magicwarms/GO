package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// var a bool = true
	// var b bool = false

	//alternate set var
	a := true
	b := false
	fmt.Println("a itu", a, "dan b itu", b)

	c := a && b
	fmt.Println("kalau a nya true DAN b nya false, c nya jadi apa hayo", c)

	d := a || b
	fmt.Println("kalau a nya true ATAU b nya false, d nya jadi apa hayo", d)

	e := 65
	fmt.Printf("type nya apa si e ini %T dan size nya adalah %d\n", e, unsafe.Sizeof(a))

	f, g := 14.08, 8.30
	fmt.Printf("type nya apa si f ini %T dan type si g apa ya %d\n", f, unsafe.Sizeof(g))

	tambah := f + g
	kurang := f - g
	fmt.Println("penjumlahan:", tambah, "pengurangan:", kurang)

	var positiveNumber uint8 = 234
	var negativeNumber int16 = -234

	fmt.Println("ini positif number", positiveNumber, "ini negatif number", negativeNumber)

	//merge string
	first := "Andhana"
	last := "Utama"

	merge := first + " " + last

	fmt.Println(merge)

	//type conversion
	h := 2
	i := 14.80
	fmt.Println("jadi nya ketika di convert", int(i))
	total := h + int(i) //this convert start
	fmt.Println(total)

	//explicit conversion
	j := 10
	var k float64 = float64(j)
	fmt.Printf("type nya apa si k ini %T dan size si k berapa ya %d\n", k, unsafe.Sizeof(k))
	fmt.Println("k adalah: ", k)
}
