package main

import "fmt"

// func orderSomeFood(menu string) {
// 	defer fmt.Println("Terima kasih, mohon menunggu!")
// 	if menu == "pizza" {
// 		fmt.Println("Pesanan spesial nih:", menu)
// 		fmt.Println("Anda akan mendapatkan promo gratis.")
// 		// os.Exit(1)
// 		return
// 	}
// 	fmt.Println("Pesanan anda tu:", menu)
// }

// func loggingApp() {
// 	fmt.Println("EKSEKUSI DI AKHIR")
// }

// func runApp() {
// 	fmt.Println("EKSEKUSI DI AWAL")
// }

func endApp() {
	// assign recover function, jadi aplikasi tetep berjalan gak berhenti
	message := recover()
	if message != nil {
		fmt.Println("Error with message", message)
	}
	fmt.Println("App process ended")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR WAK!")
	}
	fmt.Println("App process working")
}

func main() {
	// bahwa defer digunakan untuk mengakhirkan eksekusi baris kode dalam skope blok fungsi.
	// Defer bisa ditempatkan di mana saja, awal maupun akhir blok.
	// Tetapi tidak mempengaruhi kapan waktu dieksekusinya, akan selalu dieksekusi di akhir.
	// defer fmt.Println("Aku di jalankan di akhir wak")
	// fmt.Println("Ngardak!")

	// orderSomeFood("pizza")
	// orderSomeFood("boba")

	// defer loggingApp()

	// runApp()

	// panic function di golang
	runApp(true)
	fmt.Println("line code ini tetep berjalan jika kita pakai recover function")

}
