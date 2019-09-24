package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "А роза упала на лапу Азора"
	fmt.Println("Word =", str, " is palindrom =", Palindrom(str))
	fmt.Println("Word =", str, " is palindrom =", PalindromNotRecurs(str))
}

//Palindrom aa
func Palindrom(str string) bool {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ToLower(str)
	return Pali(str)
}

//Pali aa
func Pali(str string) bool {
	if utf8.RuneCountInString(str) <= 1 {
		return true
	}
	first, sizeofFirst := utf8.DecodeRuneInString(str)
	second, sizeofsecond := utf8.DecodeLastRuneInString(str)
	if first != second {
		return false
	}
	return Pali(str[sizeofFirst : len(str)-sizeofsecond])
}

//PalindromNotRecurs aa
func PalindromNotRecurs(str string) bool {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ToLower(str)
	return PaliNotRecurs(str)
}

//PaliNotRecurs aa
func PaliNotRecurs(str string) bool {
	if utf8.RuneCountInString(str) <= 1 {
		return true
	}

	for len(str) > 0 {
		first, _ := utf8.DecodeRuneInString(str)
		second, _ := utf8.DecodeLastRuneInString(str)
		if first != second {
			return false
		}
		str = strings.TrimPrefix(str, string(first))
		str = strings.TrimSuffix(str, string(second))

	}
	return true
}
