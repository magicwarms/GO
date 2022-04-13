package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
	numbOfProcess := flag.Int("process", 1, "Put your number of process")
	flag.Parse()

	fmt.Println(*host, *username, *password, *numbOfProcess)
}
