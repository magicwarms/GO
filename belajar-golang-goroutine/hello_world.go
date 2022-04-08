package belajar_golang_goroutine

import "log"

func HelloWorld() {
	log.Println("Hello World yang fana ini")
}

func DisplayNumber(number int) {
	log.Println("Nomor ke: ", number)
}

// ini channel in
// untuk mengirim/memasukkan data
func GiveMeResponse(channel chan<- string) {
	log.Println("Function 'GiveMeResponse' called")

	channel <- "Andhana Utama's Channel Wak hehe"

	log.Println("Function 'GiveMeResponse' ended")
}

// ini channel out
// untuk menerima data
func GiveMeResult(channel <-chan string) {
	log.Println("Function 'GiveMeResult' called")
	data := <-channel

	log.Println(data)

	log.Println("Function 'GiveMeResult' ended")
}
