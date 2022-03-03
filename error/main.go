package main

import (
	"errors"
	"fmt"
)

// func validate(input string) (bool, error) {
// 	if strings.TrimSpace(input) == "" {
// 		return false, errors.New("input cannot be empty")
// 	}
// 	return true, nil
// }

// func catch() {
// 	if r := recover(); r != nil {
// 		fmt.Println("Error occured", r)
// 	} else {
// 		fmt.Println("Application running perfectly")
// 	}
// }

func main() {
	// var input string
	// fmt.Print("Type some number: ")
	// fmt.Scanln(&input)

	// var number int
	// var err error
	// number, err = strconv.Atoi(input)
	// if err == nil {
	// 	fmt.Println(number, "is number")
	// } else {
	// 	fmt.Println(input, "is not number")
	// 	fmt.Println(err.Error())
	// }
	// defer catch()
	// var name string
	// fmt.Print("Type your name: ")
	// fmt.Scanln(&name)

	// if valid, err := validate(name); valid {
	// 	fmt.Println("halo", name)
	// } else {
	// 	panic(err.Error())
	// }

	// contohError := errors.New("opps, error wak")
	// fmt.Println(contohError)

	hasil, err := pembagian(100, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(hasil)
	}

}

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0")
	}
	return nilai / pembagi, nil
}
