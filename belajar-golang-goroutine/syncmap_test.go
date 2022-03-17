package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func addToMap(group *sync.WaitGroup, data *sync.Map, value int) {
	defer group.Done()

	data.Store(value, value)
}

func TestSyncMap(t *testing.T) {
	group := &sync.WaitGroup{}
	data := &sync.Map{}

	for i := 1; i <= 10; i++ {
		group.Add(1)
		go addToMap(group, data, i)
	}

	group.Wait()
	// pakai sync.Map untuk iterasi mendapatkan hasil dari goroutine, jangan golang map biasa,
	// karna udah dihandle untuk mencegah race condition
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})

}
