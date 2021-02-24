package main

import "fmt"
import "time"

func main() {
	timeNow := time.Now()
	fmt.Println(timeNow)
	fmt.Println(timeNow.Location().String())
}
