package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				// jadi mutex ini semacam mengunci blok sebuah program/proses
				// mutex ini hanya memperbolehkan 1 goroutine yang lock
				mutex.Lock() // ini untuk handle race condition di golang
				x = x + 1
				mutex.Unlock() // dan wajib juga di unlock juga setelah locked
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Total counter: ", x)
}
