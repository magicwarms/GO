package main

// parameter dan return type adalah optional di go lang function
// func calculateBill(qty int, price int) int {
// 	total := qty * price

// 	return total
// }

// other way to create function in golang
// kalo type param nya sama gak usah dibikin kayak param function diatas
// func calculateBill(qty, price int) int {
// 	// total := qty * price
// 	return qty * price
// }

// example multiple return value result from function
// func rectProps(length, width float64) (float64, float64) {
// 	area := length * width
// 	perimeter := (length + width) * 2
// 	return area, perimeter
// }

// other way to create and function declaration as return values
// func rectProps(length, width float64) (area, perimeter float64) {
// 	area = length * width
// 	perimeter = (length + width) * 2
// 	return //no explicit return value
// }

// func printMessage(message string, arr []string) {
// 	nameString := strings.Join(arr, " and ")
// 	fmt.Println(message, nameString)
// }

// func divideValue(v1, v2 float64) float64 {
// 	if v1 == 0 {
// 		fmt.Println("Nilai pembagian tidak boleh NOL!")
// 		return 0
// 	}
// 	res := v1 / v2
// 	return res
// }

// func sayHello(angkaKe int) {
// 	fmt.Println("Hallo Wak!", angkaKe)
// }

// func getHello(name string) string {
// 	println(name)
// 	if name == "" {
// 		return "Hallo wak"
// 	} else {
// 		return "Hallo " + name
// 	}
// }

// type dibawah adalah type declarations, untuk mempersingkat membuat function as parameter
type Filter func(string) string
type Blacklist func(string) bool

// sebelum nya seperti ini
// func sayHelloWithFilter(name string, spamFilter func(string) string) {
// jadi sperti ini
// func sayHelloWithFilter(name string, spamFilter Filter) {
// 	nameFiltered := spamFilter(name)
// 	fmt.Println("hello", nameFiltered)
// }

// func spamFilter(name string) string {
// 	if name == "Anjing" {
// 		return "..."
// 	} else {
// 		return name
// 	}
// }

// func registerUser(name string, blacklist Blacklist) string {
// 	if blacklist(name) {
// 		return "You are blocked " + name
// 	} else {
// 		return "Hi user " + name
// 	}
// }

// func factorialRecursive(value int) int {
// 	if value == 1 {
// 		return value
// 	}
// 	return value * factorialRecursive(value-1)
// }

func main() {
	// qty, price := 2, 2000
	// fmt.Println("Total pembayaran kamu:", calculateBill(qty, price))

	// multiple return from function
	// area, perimeter := rectProps(10.8, 5.6)
	// fmt.Printf("Area %f Perimeter %f", area, perimeter)
	// bagiDulu := divideValue(7, 2)
	// fmt.Println(bagiDulu)
	// fmt.Println(math.Pi)
	// names := []string{"Andhana", "Rani"}
	// printMessage("halo", names)
	//kalo mau kembalikan hanya 1 return value dari function pakai _ (underscore)
	// area, _ := rectProps(10.8, 5.6) // perimeter variable is discarded
	// fmt.Printf("Area %f ", area)

	// for i := 1; i <= 10; i++ {
	// 	sayHello(i)
	// }

	// result := getHello("")
	// fmt.Println(result)

	// function as parameter
	// sayHelloWithFilter("Andhana", spamFilter)
	// sayHelloWithFilter("Anjing", spamFilter)

	// anonymous function
	// blacklist := func(name string) bool {
	// 	return name == "admin"
	// }
	// registerUserData := registerUser("andhana", blacklist)
	// registerUserDataBlack := registerUser("admin", blacklist)

	// fmt.Println(registerUserData, registerUserDataBlack)

	// recursive function
	// fmt.Println(factorialRecursive(5))
}
