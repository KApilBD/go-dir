package data

import "fmt"

var Countries [10]string
var Slice []int
var Codes map[int]string

func init() {
	Countries[0] = "Argentina"
	Countries[1] = "India"
	Countries[8] = "Japan"

	qty := len(Countries)

	Prices := [2]int{100, 150}

	//Slices
	names := []string{"Kapil", "Atrayee"}
	namesw := append(names, "Mammy")

	fmt.Println(len(names), Prices, namesw)

	wellIKnownPorts := map[string]int{"http": 80, "https": 443}
	fmt.Println(wellIKnownPorts)

	fmt.Println("Countries Save", Countries, qty)
}
