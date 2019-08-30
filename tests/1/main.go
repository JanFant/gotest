package main

import "fmt"

func main() {
	s := []string{"A", "B", "C", "D", "E", "F", "G"}
	x := RemoveStringSlice(s, 0, 2)      // Удалит s[:2] из начала
	y := RemoveStringSlice(s, 1, 5)      // Удалит s[1:5] из середины
	z := RemoveStringSlice(s, 4, len(s)) // Удалит s[4:] в конце
	fmt.Printf("%v\n%v\n%v\n%v\n", s, x, y, z)
}

func RemoveStringSlice(slice []string, start, end int) []string {
	return append(slice[:start], slice[end:]...)
}

func RemoveStringSliceCopy(slice []string, start, end int) []string {
	result := make([]string, len(slice)-(end-start))
	at := copy(result, slice[:start])
	copy(result[at:], slice[end:])
	return result
}
