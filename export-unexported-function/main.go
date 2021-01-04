package main

import (
	// pemanfaatan alias pada saat import package
	libraryBiasa "export-unexported-function/library"
	libraryStruct "export-unexported-function/library2"
	"fmt"
)

func main() {
	libraryBiasa.SayHello("Yuna")

	var s1 = libraryStruct.Student{Name: "Lala", Grade: 2}

	fmt.Println("Nama nya:", s1.Name)
	fmt.Println("Kelas nya:", s1.Grade)
}
