package main

import (
	"fmt"
	"runtime"
)

func sendMsg(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("Kirim %d", i)
	}
	close(ch)
}

func receiveMsg(ch <-chan string) {
	for msg := range ch {
		fmt.Println("Terima:", msg)
	}
}

func main() {
	numOfCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numOfCores)

	var msg = make(chan string)
	go sendMsg(msg)
	receiveMsg(msg)
}
