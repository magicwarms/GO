package library

import (
	"fmt"
)

var StudentInit = struct {
	Name  string
	Grade int
}{}

func init() {
	StudentInit.Name = "Sempardak"
	StudentInit.Grade = 10
	fmt.Println("Library imported niih ------>")
}

// SayHello <--- harus kayak gini kalo mau comment EXPORTED function
// function ini EXPORTED(bisa di export di luar file ini) karna diawali dengan huruf besar
func SayHello(name string) {
	fmt.Println("Hello sayang")
	// cara nya agar bisa dipanggil di SayHello function,
	// panggil function UNEXPORTED tadi di function ini
	introduce(name)
}

// function ini UNEXPORTED(gak bisa di export di luar file ini) karna diawali dengan huruf kecil
func introduce(name string) {
	fmt.Println("Nama saya:", name)
}
