package belajar_golang_goroutine

import (
	"fmt"
)

func HelloWorld() {
	fmt.Println("Hello World yang fana ini")
}

func DisplayNumber(number int) {
	fmt.Println("Nomor ke: ", number)
}

// ini channel in
// untuk mengirim/memasukkan data
func GiveMeResponse(channel chan<- string) {
	fmt.Println("Function 'GiveMeResponse' called")

	channel <- "Andhana Utama Channel's Wak hehe"

	fmt.Println("Function 'GiveMeResponse' ended")
}

// ini channel out
// untuk menerima data
func GiveMeResult(channel <-chan string) {
	fmt.Println("Function 'GiveMeResult' called")
	data := <-channel

	fmt.Println(data)

	fmt.Println("Function 'GiveMeResult' ended")
}
