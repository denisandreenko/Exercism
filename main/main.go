package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "3-598-21508-9"
	input = strings.ReplaceAll(input, "-", "")
	fmt.Println(input[0])
	fmt.Println(rune(input[0]))
	fmt.Println(string(rune(input[0])))
}