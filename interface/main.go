package main

import (
	"fmt"
	"math"
)

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	// luas() float64
	// keliling() float64
	hitung2d
	hitung3d
}

// type lingkaran struct {
// 	diameter float64
// }

// func (l lingkaran) jariJari() float64 {
// 	return l.diameter / 2
// }

// func (l lingkaran) luas() float64 {
// 	return math.Pi * math.Pow(l.jariJari(), 2)
// }

// func (l lingkaran) keliling() float64 {
// 	return math.Pi * l.diameter
// }

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func (p persegi) volume() float64 {
	return p.sisi * p.sisi * p.sisi
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func main() {
	var bangunDatar hitung

	fmt.Println("=== PERSEGI ===")
	bangunDatar = persegi{10}
	fmt.Println("luas 		:", bangunDatar.luas())
	fmt.Println("keliling 	:", bangunDatar.keliling())
	fmt.Println("volume 	:", bangunDatar.volume())

	// fmt.Println("=== LINGKARAN ===")
	// bangunDatar = lingkaran{14}
	// fmt.Println("luas 		:", bangunDatar.luas())
	// fmt.Println("keliling 	:", bangunDatar.keliling())
	// fmt.Println("jari-jari 	:", bangunDatar.(lingkaran).jariJari())

	var bangunRuang hitung = &kubus{4}

	fmt.Println("=== KUBUS ===")
	fmt.Println("luas      :", bangunRuang.luas())
	fmt.Println("keliling  :", bangunRuang.keliling())
	fmt.Println("volume    :", bangunRuang.volume())
}
