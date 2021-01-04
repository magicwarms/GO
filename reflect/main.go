package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Reflection adalah teknik untuk inspeksi variabel,
	// mengambil informasi dari variabel tersebut atau bahkan memanipulasinya.
	var number = 14
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variable ini:", reflectValue.Type())
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variable ini:", reflectValue.Int())
	}
}
