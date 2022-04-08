package belajar_golang_goroutine

import (
	"log"
	"sync"
	"testing"
	"time"
)

// timer ini semacam delay job gitu
func TestTimer(t *testing.T) {
	// ini cara manual timer
	timer := time.NewTimer(5 * time.Second)
	log.Println(time.Now())

	time := <-timer.C

	log.Println(time)
}

func TestAfter(t *testing.T) {
	// ini cara otomatis timer
	channel := time.After(5 * time.Second)
	log.Println(time.Now())

	time := <-channel

	log.Println(time)
}

// ini function untuk memproses function dengan menggunakan delay
func TestAfterFunc(t *testing.T) {
	log.Println("Proses dimulai")
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(2*time.Second, func() {
		log.Println("DI PROSES SETELAH 2 DETIK")
		group.Done()
	})
	// jalanin code dibawah dulu dia, baru diatas
	log.Println()

	group.Wait()
	log.Println("Proses selesai")
}
