package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)
	// Cara pembuatan channel yaitu dengan menuliskan keyword 'make'
	// dengan isi keyword 'chan' diikuti dengan tipe data channel yang diinginkan.
	// var messages = make(chan string) // <-- ini buat channel baru nya

	// Buffered Channel
	// Pada kode di dibawah, parameter kedua fungsi make() adalah representasi jumlah buffer.
	// Perlu diperhatikan bahwa nilai
	// buffered channel dimulai dari 0. Ketika nilainya adalah 2 brarti jumlah buffer maksimal ada 3.
	var messages = make(chan int, 3)

	// var sayHelloTo = func(who string) {
	// 	var data = fmt.Sprintf("hello %s", who)
	// 	messages <- data
	// }

	// go sayHelloTo("john wick")
	// go sayHelloTo("ethan hunt")
	// go sayHelloTo("jason bourne")

	// var message1 = <-messages
	// fmt.Println(message1)

	// var message2 = <-messages
	// fmt.Println(message2)

	// var message3 = <-messages
	// fmt.Println(message3)

	// penjelasan kode dibawah
	// jadi tersedia data array, looping dulu untuk assign name ke var channel messages diatas
	// var names = []string{"1 john wick", "2 ethan hunt", "3 jason bourne"}
	// for _, val := range names {
	// 	var assignName = func(who string) {
	// 		var data = fmt.Sprintf(who)
	// 		messages <- data
	// 	}
	// 	// panggil function closure diatas lalu assign ke variable messages
	// 	go assignName(val)
	// }
	// lalu looping data nya
	// jadi yang aku lakuin ini adalah channel sebagai tipe data parameter
	// for i := 0; i < len(names); i++ {
	// 	sayHelloTo(messages)
	// }

	for i := 1; i <= 5; i++ {
		fmt.Println("send data:", i)
		messages <- i
		go receiveData(messages)
	}
	// // dikasih gini supaya proses nya bener2 selesai
	// var input string
	// fmt.Scanln(&input)
}

func receiveData(receive chan int) {
	fmt.Println("receive data:", <-receive)
}

func sayHelloTo(who chan string) {
	fmt.Println("hello wak", <-who)
}
