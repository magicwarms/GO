package main

import (
	"fmt"
	"runtime"
)

func massivePrint(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	// Fungsi ini digunakan untuk menentukan
	// jumlah core atau processor yang digunakan dalam eksekusi program.
	numOfCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numOfCores)
	// Pembuatan goroutine baru ditandai dengan keyword go.
	go massivePrint(5, "mamak")
	massivePrint(5, "bapak")

	var input string
	// fungsi Scanln mengakibatkan proses jalannya aplikasi berhenti di baris itu (blocking)
	// hingga user menekan tombol enter.
	fmt.Scanln(&input)

	fmt.Println(input)
}
