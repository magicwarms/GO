package main

import (
	"fmt"
	"runtime"
)

func main() {
	numOfCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numOfCores)
	// Cara pembuatan channel yaitu dengan menuliskan keyword 'make'
	// dengan isi keyword 'chan' diikuti dengan tipe data channel yang diinginkan.
	// var messages = make(chan string) // <-- ini buat channel baru nya

	// Buffered Channel
	// Pada kode di dibawah, parameter kedua fungsi make() adalah representasi jumlah buffer.
	// Perlu diperhatikan bahwa nilai
	// buffered channel dimulai dari 0. Ketika nilainya adalah 2 brarti jumlah buffer maksimal ada 3.
	var messages = make(chan string, 1)

	// dibawah ini function closure wak!
	// var sayHelloTo = func(who int) {
	// 	var data = who
	// 	nombor <- data
	// }

	// go sayHelloTo(1)
	// go sayHelloTo(2)
	// go sayHelloTo(3)

	// message1 := <-nombor
	// fmt.Println(message1)

	// message2 := <-nombor
	// fmt.Println(message2)

	// message3 := <-nombor
	// fmt.Println(message3)

	// penjelasan kode dibawah
	// jadi tersedia data array, looping dulu untuk assign name ke var channel messages diatas
	var names = []string{"1 john wick", "2 ethan hunt", "3 jason bourne"}
	for _, val := range names {
		var assignName = func(who string) {
			var data = fmt.Sprint(who)
			messages <- data
		}
		// panggil function closure diatas lalu assign ke variable messages
		go assignName(val)
	}
	// lalu looping data nya
	// jadi yang aku lakuin ini adalah channel sebagai tipe data parameter
	for i := 0; i < len(names); i++ {
		sayHelloTo(messages)
	}
	// nombor := make(chan int)
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println("prepare send data:", i)
	// 	nombor <- i
	// 	receiveData(nombor)
	// 	fmt.Println("data sent: ", i)
	// }

	// dikasih gini supaya proses nya bener2 selesai
	// var input string
	// fmt.Scanln(&input)
}

// func receiveData(receive chan int) {
// 	fmt.Println("received data:", <-receive)
// }

func sayHelloTo(who chan string) {
	fmt.Println("hello wak", <-who)
}
