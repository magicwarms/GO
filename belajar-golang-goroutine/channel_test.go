package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	// pakai defer aja lebih aman wak
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		// kalau gak ada kirim dan cuman ada nerima, bakal hang server nya/process
		// kirim data
		channel <- "Andhana Utama"
		fmt.Println("Selesai Mengirim data ke Channel")
	}()

	// kalo gak ada yang ambil data di channel dia bakal ngeblock terus, diem terus dan terjadi error
	// terima data
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

	// cara kirim data ke channel
	// channel <- "Andhana"

	// cara ambil data dari channel
	// data := <-channel
	// fmt.Println(data)

}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	go GiveMeResult(channel)

	time.Sleep(2 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	// tambahin buffered disini contoh: 4 channel biar channel nya gak blocking
	channel := make(chan string, 4)
	defer close(channel)

	channel <- "Andhana"
	channel <- "Mahrani"
	channel <- "Yuna"
	channel <- "Sorra"

	// print result nya 1 per 1
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	// kalau keluarin data nya yg gak pernah terkirim ke channel, bakal error deadlock
	// fmt.Println(<-channel)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Iterasi ke berapa ini " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counterWak := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Dari dari channel 1", data)
			counterWak++
		case data := <-channel2:
			fmt.Println("Dari dari channel 2", data)
			counterWak++
		default: //bakal jalan pertama kali pada saat select range
			fmt.Println("Menunggu data")
		}

		if counterWak == 2 {
			break
		}
	}
}
