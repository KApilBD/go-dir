package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)

		time.Sleep(800 * time.Millisecond)
	}

	channel <- "Done bro!"
}

func main() {
	// var channel chan string
	channel := make(chan string)
	go printMessage("Go is great !!!!", channel)
	// go printMessage("FM Rocks \\m/")
	res := <-channel
	fmt.Println(res)
}

// func main() {
// 	go printMessage("Go is great !!!!")
// 	go printMessage("FM Rocks \\m/")
// 	go printMessage("We miss classes ;(")
// 	fmt.Println("Testttttt")
// 	fmt.Println("Testttttt 222222")

// }
