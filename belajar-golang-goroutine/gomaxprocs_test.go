package belajar_golang_goroutine

import (
	"log"
	"runtime"
	"testing"
)

func TestGoMaxProcs(t *testing.T) {
	totalCPU := runtime.NumCPU()
	log.Println("Total CPU:", totalCPU)
	// ini untuk mengubah thread
	// tak usah diubah gomaxprocs karna default nya sudah optimal
	// runtime.GOMAXPROCS(20) // <-- tak usah diubah-ubah ini
	totalThread := runtime.GOMAXPROCS(-1)
	log.Println("Total Thread:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	log.Println("total Goroutine:", totalGoroutine)

}
