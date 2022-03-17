package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

// sync.WaitGroup ini nih kayak tunggu process semua goroutine selesai

func RunAsynchronous(group *sync.WaitGroup, counter int) {
	// wajib di kasih group.Done()
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello", counter)
	// time.Sleep(1 * time.Second)

}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group, i)
	}

	group.Wait()
	fmt.Println("COMPLETED")
}
