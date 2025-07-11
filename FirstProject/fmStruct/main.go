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
	fmt.Println(jishnu.Print())

	goCourse := data.Course{Id: 2, Name: "Go Fundamentals", Instructor: jishnu}
	fmt.Printf("%v", goCourse)

	swiftWS := data.NewWorkshop("Swift with iOS", jishnu)
	fmt.Printf("%v", swiftWS)

	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swiftWS

	for i, course := range courses {
		fmt.Println(i, course)
	}

}
