package main

func main() {
	// var month map[string]int
	// month = map[string]int{}
	// fmt.Println("ini month", len(month))
	// Zero value dari map adalah nil, maka tiap variabel bertipe map harus di-inisialisasi secara explisit nilai awalnya (agar tidak nil).

	// var data map[string]int
	// data["one"] = 1
	// akan muncul error!

	// data = map[string]int{}
	// data["one"] = 1
	// tidak ada error

	// month["January"] = 01
	// month["November"] = 11
	// // Pengisian data pada map bersifat overwrite
	// month["November"] = 12
	// fmt.Println("Bulan pertama apa:", month["January"])
	// fmt.Println("Bulan kesebelas apa:", month["November"])

	// cara inisialisasi map
	// cara vertikal

	// var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// cara horizontal
	// Khusus deklarasi dengan gaya horizontal, tanda koma perlu dituliskan setelah item terakhir.
	// var chicken2 = map[string]int{
	// 	"januari":  50,
	// 	"februari": 40,
	// }

	// Variabel map bisa di-inisialisasi dengan tanpa nilai awal
	// caranya menggunakan tanda kurung kurawal
	// atau
	// bisa juga dengan menggunakan keyword make dan new. Contohnya bisa dilihat pada kode berikut. Ketiga cara di bawah ini intinya adalah sama.
	// var chicken3 = map[string]int{}
	// var chicken4 = make(map[string]int)
	// Khusus inisialisasi data menggunakan keyword new, yang dihasilkan adalah data pointer. Untuk mengambil nilai aslinya bisa dengan menggunakan tanda asterisk (*)
	// var chicken5 = *new(map[string]int)

	// ini iterasi Map menggunakan for - range
	// var chicken = map[string]int{
	// 	"januari":  50,
	// 	"februari": 40,
	// 	"maret":    34,
	// 	"april":    67,
	// }

	// for key, val := range chicken {
	// 	fmt.Println(key, "  \t:", val)
	// }
	// fmt.Println(len(chicken)) // 2
	// fmt.Println(chicken)

	// Fungsi delete() digunakan untuk menghapus item dengan key tertentu pada variabel map.
	// delete(chicken, "januari")

	// fmt.Println(len(chicken)) // 1
	// fmt.Println(chicken)

	// Deteksi Keberadaan Item Dengan Key Tertentu
	// var value, isExist = chicken["januari"]
	// if isExist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("item is not exists")
	// }

	// Kombinasi Slice & Map
	// var chickens = []map[string]string{
	// 	map[string]string{"name": "chicken blue", "gender": "male"},
	// 	map[string]string{"name": "chicken red", "gender": "male"},
	// 	map[string]string{"name": "chicken yellow", "gender": "female"},
	// }
	// Jika menggunakan versi go terbaru, cara deklarasi slice-map bisa dipersingkat, tipe tiap elemen tidak wajib untuk dituliskan
	// var chickens = []map[string]string{
	// 	{"name": "chicken blue", "gender": "male", "address": "Tiban sih"},
	// 	{"name": "chicken red", "gender": "male", "address": "Tiban Loh"},
	// 	{"name": "chicken yellow", "gender": "female"},
	// }
	// for _, chicken := range chickens {
	// 	fmt.Println(chicken["gender"], chicken["name"], chicken["address"])
	// }
}
