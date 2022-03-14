package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Andhana",
		Balance: 10000,
	}
	user2 := UserBalance{
		Name:    "Mahrani",
		Balance: 1000,
	}

	go Transfer(&user1, &user2, 1000)
	go Transfer(&user2, &user1, 2000)

	time.Sleep(10 * time.Second)

	fmt.Println("User: ", user1.Name, ", Balance: ", user1.Balance)
	fmt.Println("User: ", user2.Name, ", Balance: ", user2.Balance)

}
