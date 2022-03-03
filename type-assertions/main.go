package main

import "fmt"

func random() interface{} {
	return "OK SEMPARDAK"
}

func main() {
	// type assertion hanya merubah type data nya saja biasa nya sering kali digunakan ketika bertemu dengan data interface kosong

	var result interface{} = random()
	var resultString string = result.(string)
	fmt.Println(resultString)
}
