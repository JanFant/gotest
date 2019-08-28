package main

import (
	"fmt"
	"unicode"
)

func main() {
	//162 c
	fmt.Println(IsHexDigit('8'), IsHexDigit('x'), IsHexDigit('X'), IsHexDigit('b'), IsHexDigit('B'))
}

func IsHexDigit(char rune) bool {
	return unicode.Is(unicode.ASCII_Hex_Digit, char)
}
