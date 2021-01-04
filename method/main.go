package main

import (
	"fmt"
	"strings"
)

type student struct {
	name   string
	gender int
}

func (s student) sayHello() {
	fmt.Println("hello wak", s.name)
}

func (s student) getLastName(i int) string {
	return strings.Split(s.name, " ")[i]
}

func main() {
	// Method adalah fungsi yang menempel pada type (bisa struct atau tipe data lainnya). Method bisa diakses lewat variabel objek.
	s1 := student{name: "Andhana Utama", gender: 28}

	s1.sayHello()

	lastName := s1.getLastName(1)

	fmt.Println("ini nama terakhir elo kan:", lastName)
}
