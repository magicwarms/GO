package main

type Animal struct {
	Name string
}

type Person struct {
	Fullname string
}

type HasName interface {
	GetName() string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (person Person) GetName() string {
	return person.Fullname
}

func sayHello(hasName HasName) string {
	return "Hello wak " + hasName.GetName() + "!"
}

func main() {
	animalName := Animal{Name: "Perkutut"}

	println(sayHello(animalName))

	personFullname := Person{Fullname: "Naura Yuna Utama"}

	println(sayHello(personFullname))
}
