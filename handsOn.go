package main

import (
	"fmt"
	"strings"
)

func capitalize(sentence string) string {
	result := ""
	for index, char := range sentence {
		if index == 0 && char != ' ' {
			result += strings.ToUpper(string(char))
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(capitalize("hello world!"))
}
