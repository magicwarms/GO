package main

import "fmt"

func getAverage(numbers []int, ch1 chan float64) {
	var sum = 0
	for _, val := range numbers {
		sum += val
	}
	ch1 <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch2 chan int) {
	var max = numbers[0]
	for _, val := range numbers {
		if max < val {
			max = val
		}
	}
	ch2 <- max
}

func main() {
	var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println("numbers :", numbers)

	ch1 := make(chan float64)
	go getAverage(numbers, ch1)

	ch2 := make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Max \t: %d \n", max)
		}
	}
	// avg := <-ch1
	// fmt.Printf("Avg \t: %.2f \n", avg)
	// max := <-ch2
	// fmt.Printf("Max \t: %d \n", max)

}
