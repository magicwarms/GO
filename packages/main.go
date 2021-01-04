package main

import (
	"fmt"
	"log"
	"packages/simpleinterest"
)

var principal, rateOfInterest, timeDuration = -5000.0, 10.0, 1.0

// Ini dipakai pada saat contoh function Calculate lagi in development tapi kita udah pakai function tersebut
// var _ = simpleinterest.Calculate

/*
* init function to check if principal, rateOfInterest and timeDuration are greater than zero
 */
func init() {
	println("Main package initialized")
	if principal < 0 {
		log.Fatal("Principal is less than zero")
	}
	if rateOfInterest < 0 {
		log.Fatal("Rate of interest is less than zero")
	}
	if timeDuration < 0 {
		log.Fatal("Duration is less than zero")
	}
}

func main() {
	fmt.Println("Simple kalkulasi matematika aja wak - Simple Interest")

	totalInterest := simpleinterest.Calculate(principal, rateOfInterest, timeDuration)

	fmt.Println("Simple interest is", totalInterest)
}
