package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) ChangeAmount(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("User 1 Lock: ", user1.Name)
	user1.ChangeAmount(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("User 2 Lock: ", user2.Name)
	user2.ChangeAmount(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}
