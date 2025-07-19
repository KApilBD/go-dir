package main

import (
	"fmt"
)

// struct
type Person struct {
	Name string
	Age  string
}

func main() {
	//variables
	var name string = "Kapil"
	fmt.Printf("%s!!! Hello from Go\n", name)

	age := 27
	fmt.Printf("My age is : %d \n", age)

	var city string
	city = "Bangalore"
	fmt.Printf("I live in %s. \n", city)

	var (
		isEmployed bool   = true
		salary     int    = 50000
		position   string = "developer"
	)
	fmt.Printf("I am working %t, as %s with %d compansation\n", isEmployed, position, salary)

	const pi = 3.24

	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
	)
	fmt.Printf("Monday: %d, Tuesday %d, Wednesday %d \n", Monday, Tuesday, Wednesday)

	//ENUM
	const (
		Jan int = iota + 1
		Feb
		Mar
		Apr
	)
	fmt.Printf("jan: %d, feb %d, mar %d, apr %d \n", Jan, Feb, Mar, Apr)

	fmt.Println(add(3, 4))

	sum, substract := addandsubstract(12, 4)
	fmt.Println(sum, substract)

	//Arrays

	numbers := [5]int{10, 20, 30, 40, 50}

	numbersATInit := [...]int{10, 20, 30, 40, 50, 60}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	fmt.Println(numbersATInit)

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix)

	//Slice

	allNumbers := numbers[:]
	firstThree := numbers[0:3]

	fmt.Println("slice 1: ", allNumbers)
	fmt.Println("slice 2:", firstThree)

	fruits := []string{"apple", "banana", "orange"}
	fruits = append(fruits, "kiwi")

	fmt.Println(fruits)

	fruits = append(fruits, "mango", "cherry")

	fmt.Println(fruits)
	fruits = append(fruits, "mango")
	fmt.Println(fruits)

	for index, value := range fruits {
		fmt.Println(index, value)
	}

	//Maps

	capitalCities := map[string]string{
		"USA":   "W DC",
		"India": "Delhi",
		"UK":    "London",
	}

	fmt.Println(capitalCities["India"])
	capital, exists := capitalCities["Germany"]
	if exists {
		fmt.Println(" this the capital : ", capital)
	} else {
		fmt.Println("not there")
	}

	delete(capitalCities, "UK")

	fmt.Println(capitalCities)

	// Structs
	person := Person{Name: "Kapil", Age: "65"}

	fmt.Printf("The Person: %+v\n", person)

	// anonymus struct
	atrist := struct {
		name string
		id   int
	}{
		name: "Atrayee",
		id:   1,
	}

	fmt.Printf("The atrist: %+v\n", atrist)

	fmt.Println("name before:", person.Name)
	person.modifyPersonName("Kapilraj Baraskar")
	fmt.Println("name after:", person.Name)

}

func add(a int, b int) int {
	return a + b
}

func addandsubstract(a, b int) (int, int) {
	return a + b, a - b
}

func (p *Person) modifyPersonName(name string) {
	p.Name = name
	fmt.Println("inside scooe: name:", p.Name)
}
