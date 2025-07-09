package main

import (
	"fm/go/server/data"
	"fmt"
)

func main() {
	fmt.Println("Structures ::::")
	jishnu := data.Instructor{}
	jishnu.Id = 1
	jishnu.FirstName = "Jishnu"
	jishnu.LastName = "Bhattacharya"
	fmt.Println(jishnu)

	fmt.Println(jishnu.Print())

	kyle := data.NewInstructor("Kyle", "Simpson")
	fmt.Println(kyle.Print())

}
