package simpleinterest

import "fmt"

/*
 * init function added
 */
func init() {
	fmt.Println("Simple interest package initialized")
}

// Calculate calculates and returns the simple interest for a principal p, rate of interest r for time duration t years
// Any variable or function which starts with a capital letter are exported names in go. Only exported functions and variables can be accessed from other packages.
func Calculate(p, r, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}
