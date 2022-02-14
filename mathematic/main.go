package main

import "fmt"

func main() {
	// ini hasil nya 20
	var result int = 10 + 10
	// ini augmented assignments
	// tambah lagi 10 tapi dengan cara augmented assignments
	// *= 10 --> kali 10 --> artinya --> result = result * 10
	// /= 10 --> bagi 10 --> artinya --> result = result / 10
	// += 10 --> tambah 10 --> artinya --> result = result + 10
	// -= 10 --> kurang 10 --> artinya --> result = result - 10
	// ini hasil nya 30
	result += 10

	var a int = 10
	var b int = 10
	var total = a * b

	fmt.Println()
	fmt.Println(result)
	fmt.Println(total)

	// unary operator
	// jadi tambah 1 dengan ditambah ++/dikurang --
	result++

	fmt.Println()
	fmt.Println(result)
}
