package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		// default value kalau kosong
		New: func() interface{} {
			return "Kosong"
		},
	}
	pool.Put("Andhana")
	pool.Put("Mahrani")
	pool.Put("Yuna")
	pool.Put("Sorra")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("selesai")
}
