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

	fmt.Println()
	c := a && b
	fmt.Println("kalau a nya true DAN b nya false, c nya jadi apa hayo", c)

	fmt.Println()
	d := a || b
	fmt.Println("kalau a nya true ATAU b nya false, d nya jadi apa hayo", d)

	fmt.Println()
	e := 65
	fmt.Printf("type nya apa si e ini %T dan size nya adalah %d\n", e, unsafe.Sizeof(a))

	fmt.Println()
	f, g := 14.08, 8.30
	fmt.Printf("type nya apa si f ini %T dan type si g apa ya %d\n", f, unsafe.Sizeof(g))

	// perbandingan string
	var name1 string = "Sorra"
	var name2 string = "Yuna"

	var startCompareName bool = name1 == name2

	fmt.Println(startCompareName)
	fmt.Println()

	// perbandingan integer
	var val1 int = 100
	var val2 int = 200

	var startCompareVal bool = val1 >= val2

	fmt.Println(startCompareVal)
	fmt.Println()
}
