package main

import (
	"log"
	"strings"
)

func main() {
	log.Println(strings.Contains("Andhana Utama", "andhana"))
	log.Println(strings.Split("Andhana Utama", " "))
	log.Println(strings.ToLower("Andhana Utama"))
	log.Println(strings.ToUpper("Andhana Utama"))
	log.Println(strings.Trim("          Andhana Utama               ", " "))
	log.Println(strings.ReplaceAll("apa tuih malas ih apa tu apa tu", "apa tu", ""))
}
