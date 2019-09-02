package main

import (
	"fmt"
	"sort"
	"strings"
)

//SortFoldedStrings aaa
func SortFoldedStrings(slice []string) {
	sort.Sort(FoldedStrings(slice))
}

// FoldedStrings aaa
type FoldedStrings []string

func (slice FoldedStrings) Len() int {
	return len(slice)
}

func (slice FoldedStrings) Less(i, j int) bool {
	return strings.ToLower(slice[i]) < strings.ToLower(slice[j])
}

func (slice FoldedStrings) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func main() {
	files := []string{"Test.conf", "util.go", "Makefile", "misc.go", "main.go"}
	fmt.Printf("%q\n", files)
	// target := "main.go"
	// for i, file := range files {
	// 	if file == target {
	// 		fmt.Printf("1found \"%s\" at files[%d]\n", file, i)
	// 		break
	// 	}
	// }

	// sort.Strings(files)
	// fmt.Printf("%q\n", files)
	// i := sort.Search(len(files), func(i int) bool { return files[i] >= target })
	// if i < len(files) && files[i] == target {
	// 	fmt.Printf("2found \"%s\" at files[%d]\n", files[i], i)
	// }
	target := "makefile"
	SortFoldedStrings(files)
	fmt.Printf("%q\n", files)
	caseInsensitiveCompare := func(i int) bool {
		return strings.ToLower(files[i]) >= target
	}
	i := sort.Search(len(files), caseInsensitiveCompare)
	if i < len(files) && strings.EqualFold(files[i], target) {
		fmt.Printf("found \"%s\" at files[%d]\n", files[i], i)
	}

}
