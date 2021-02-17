package main

import (
	"fmt"
)

type person struct {
	name string
	age  int64
}

func main() {
	// var data map[string]interface{}

	// data = map[string]interface{}{
	// 	"name":      "ethan hunt",
	// 	"grade":     2,
	// 	"breakfast": []string{"apple", "manggo", "banana"},
	// }
	// var secret interface{}

	// secret = 2
	// var number = secret.(int) * 10
	// fmt.Println(secret, "multiplied by 10 is :", number)

	// secret = []string{"apple", "manggo", "banana"}
	// var gruits = strings.Join(secret.([]string), ", ")
	// fmt.Println(gruits, "is my favorite fruits")

	// var humanPerson interface{} = person{name: "Andhana Utama", age: 28}

	// var getName = humanPerson.(person).name
	// fmt.Println((getName))

	var persons = []map[string]interface{}{
		{"name": "Wick", "age": 23},
		{"name": "Ethan", "age": 23},
		{"name": "Bourne", "age": 22},
	}

	for _, person := range persons {
		fmt.Println(person["name"], "age is", person["age"])
	}
}
