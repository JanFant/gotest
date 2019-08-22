package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var bigDigits = [][]string{
	{"  000  ", " 0   0 ", " 0   0 ", " 0   0 ", "  000  "},
	{"  1  ", " 11  ", "  1  ", "  1  ", "11111"},
	{"22222", "    2", "   2 ", " 2   ", "22222"},
	{"33333", "    3", "33333", "    3", "33333"},
	{"4   4", "4   4", "44444", "    4", "    4"},
	{"55555", "5    ", "55555", "    5", "55555"},
	{"66666", "6    ", "66666", "6   6", "66666"},
	{"77777", "   7 ", "  7  ", " 7   ", "7    "},
	{"88888", "8   8", "88888", "8   8", "88888"},
	{"99999", "9   9", "99999", "    9", "99999"},
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage:" + filepath.Base(os.Args[0]) + " <whole-number>\n")
		os.Exit(1)
	}
	stringOfDigits := os.Args[1]
	for row := range bigDigits[0] {
		line := ""
		for colums := range stringOfDigits {
			digits := stringOfDigits[colums] - '0'
			if 0 <= digits && digits <= 9 {
				line += bigDigits[digits][row] + "  "
			} else {
				log.Fatal("Invalid suqa")
			}
		}
		fmt.Println(line)
	}
}
