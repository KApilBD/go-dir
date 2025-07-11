package main

import (
	"fmt"
	"os"

	"fm.com/files/filemanagement"
)

func main() {
	path, _ := os.Getwd()
	filepath := path + "/data/text.txt"
	fmt.Println("Reading File....in wd:", path)
	content, err := filemanagement.ReadTextFile(filepath)
	if err == nil {
		fmt.Println(content)
		newContent := fmt.Sprintf("Original: %v \n Double the Original: %v%v", content, content, content)
		filemanagement.WriteToFile(filepath+".output.txt", newContent)
	} else {
		fmt.Printf("Error PANIC!!! %v", err)
	}
}
