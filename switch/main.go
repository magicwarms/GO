package main

import "fmt"

func calculateNumber(number1 int, number2 int) int {
	return number1 * number2
}

func main() {
	// finger := 4
	// fmt.Println("Finger is:", finger)
	// switch finger {
	// case 1:
	// 	fmt.Println("Thumb")
	// case 2:
	// 	fmt.Println("Index")
	// case 3:
	// 	fmt.Println("Middle")
	// case 4:
	// 	fmt.Println("Ring")
	// case 5:
	// 	fmt.Println("Pinky")
	// 	// case 5: //error nih gak boleh duplicat case nya
	// 	// 	fmt.Println("Cacat kau")
	// default: //ini default case jika kondisi diatas tak ketemu
	// 	fmt.Println("sempardak! salah bray")
	// }

	//other way to create switch case -- multiple expression in case
	// letter := "i"
	// fmt.Printf("Letter %s is a ", letter)
	// switch letter {
	// case "a", "e", "i", "o", "u": //multiple expressions in case
	// 	fmt.Println("vowel")
	// default:
	// 	fmt.Println("not a vowel")
	// }

	//other way to create switch case -- expressionless switch
	// num := 75
	// switch { // expression is omitted
	// case num >= 0 && num <= 50:
	// 	fmt.Printf("%d is greater than 0 and less than 50", num)
	// case num >= 51 && num <= 100:
	// 	fmt.Printf("%d is greater than 51 and less than 100", num)
	// case num >= 101:
	// 	fmt.Printf("%d is greater than 100", num)
	// }

	//fallthrough
	//jadi fallthrough ini akan eksekusi kondisi dibawah nya juga jika kondisi diatas terpenuhi
	// switch num := calculateNumber(); {
	// case num < 50:
	// 	if num < 0 {
	// 		break //boleh break di switch case
	// 	}
	// 	fmt.Printf("%d is lesser than 50\n", num)
	// 	fallthrough //dan ini mesti di akhir case
	// case num < 100:
	// 	fmt.Printf("%d is lesser than 100\n", num)
	// 	// fallthrough
	// case num < 200:
	// 	fmt.Printf("%d is lesser than 200", num)
	// 	// fallthrough //gak boleh taruh di akhir case wak nanti error
	// }

	//jadi fallthrough ini akan selalu eksekusi code dibawah nya walaupun kondisi nya itu false/true
	//hati-hati pakai fallthrough ini, kalo bisa hindari pakai nya
	// 	nilai := 19
	// 	switch {
	// 	case nilai < 18:
	// 		fmt.Println("Anda masih", nilai, "tahun dan anda tidak diperbolehkan masuk")
	// 	case nilai >= 18 && nilai <= 20:
	// 		fmt.Println("Anda sudah", nilai, "tahun dan anda diperbolehkan masuk")
	// 	default:
	// 		fmt.Println("Anda sudah tua cuk!")
	// 	}
	// randloop:
	// 	for {
	// 		i := rand.Intn(100)
	// 		fmt.Println(i)
	// 		switch {
	// 		case i%2 == 0:
	// 			fmt.Println("Ini genap loh", i)
	// 			break randloop
	// 		}
	// 	}

	// coba dari udemy
	numberMark := calculateNumber(6, 5)
	switch {
	case numberMark >= 30:
		fmt.Println("Nilai nya:", numberMark)
	case numberMark <= 30:
		fmt.Println("Kurang dari 30:", numberMark)
	}

}
