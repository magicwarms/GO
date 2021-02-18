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
	runtime.GOMAXPROCS(4)
	// Pembuatan goroutine baru ditandai dengan keyword go.
	go massivePrint(5, "mamak")
	massivePrint(5, "bapak")

	var input string
	fmt.Scanln(&input)

	fmt.Println(input)
}
