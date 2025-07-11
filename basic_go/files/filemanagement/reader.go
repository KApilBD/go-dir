package filemanagement

import (
	"fmt"
	"os"
)

func ReadTextFile(filename string) (string, error) {

	content, err := os.ReadFile(filename)
	fmt.Println("Content:", content, " err:", err)

	if err != nil {
		return "", err
	} else {
		return string(content), nil
	}
}
