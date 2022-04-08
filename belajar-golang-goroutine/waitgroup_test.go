package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

// sync.WaitGroup ini nih kayak tunggu process semua goroutine selesai

func RunAsynchronous(group *sync.WaitGroup, counter int) {
	// wajib di kasih group.Done()
	// kurangin 1 goroutine -1
	defer group.Done()
	//add 1 goroutine +1
	group.Add(1)

	fmt.Println("Hello", counter)
	// time.Sleep(1 * time.Second)

}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		go RunAsynchronous(group, i)
	}

	group.Wait()
	fmt.Println("COMPLETED")
}
