package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
	go HelloWorld()
	fmt.Println("WOOOY")

	time.Sleep(1 * time.Second)
}

func TestDisplayNumber(t *testing.T) {
	for i := 1; i <= 100; i++ {
		go DisplayNumber(i)
	}
	// sleep for 5 seconds
	// time.Sleep(5 * time.Second)
}
