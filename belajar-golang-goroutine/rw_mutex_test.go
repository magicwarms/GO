package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestReadWriteMutex(t *testing.T) {
	account := BankAccount{}

	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final balance: ", account.GetBalance())
}
