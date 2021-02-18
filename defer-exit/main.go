package main

import (
	"fmt"
	"os"
)

func orderSomeFood(menu string) {
	defer fmt.Println("Terima kasih, mohon menunggu!")
	if menu == "pizza" {
		fmt.Println("Pesanan spesial nih:", menu)
		fmt.Println("Anda akan mendapatkan promo gratis.")
		os.Exit(1)
		return
	}
	fmt.Println("Pesanan anda tu:", menu)
}

func main() {
	// bahwa defer digunakan untuk mengakhirkan eksekusi baris kode dalam skope blok fungsi.
	// Defer bisa ditempatkan di mana saja, awal maupun akhir blok.
	// Tetapi tidak mempengaruhi kapan waktu dieksekusinya, akan selalu dieksekusi di akhir.
	// defer fmt.Println("Aku di jalankan di akhir wak")
	// fmt.Println("Ngardak!")

	orderSomeFood("pizza")
	orderSomeFood("boba")
}
