package main

//235
import (
	"fmt"
	"log"
	"sort"
	"strings"
)

func main() {
	arrayInt := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
	irregularMatrix := [][]int{{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11},
		{12, 13, 14, 15},
		{16, 17, 18, 19, 20}}
	SforMat2d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	iniData := []string{
		"; Cut down copy of Mozilla application.ini file",
		"",
		"[App]",
		"Vendor=Mozilla",
		"Name=Iceweasel",
		"Profile=mozilla/firefox",
		"Version=3.5.16",
		"[Gecko]",
		"MinVersion=1.9.1",
		"MaxVersion=1.9.1.*",
		"[XRE]",
		"EnableProfileMigrator=0",
		"EnableExtensionManager=1",
		"[BBba]",
		"ProfileMigrator=0",
		"Manager=1",
	}
	MyparsStrint := MyParsInt(iniData)
	parsStrint := ParseIni(iniData)
	SMat2d := Mat2D(SforMat2d, 6)
	newarray := UniqueIntSlice(arrayInt)
	mar := UniqueIntMap(arrayInt)
	Slm := Smatrix(irregularMatrix)
	fmt.Println(arrayInt)
	fmt.Println(newarray)
	fmt.Println(mar)
	fmt.Println("-------------------------")
	fmt.Println(irregularMatrix)
	fmt.Println(Slm)
	fmt.Println("-------------------------")
	fmt.Println(SforMat2d)
	fmt.Println(SMat2d)
	fmt.Println("-------------------------")
	fmt.Println(iniData)
	fmt.Println(MyparsStrint)
	fmt.Println(parsStrint)
	fmt.Println("-------------------------")
	Printini(parsStrint)
	fmt.Println("-------------------------")
	PrintIni2(parsStrint)
}

// Printini  aaa
func Printini(data map[string]map[string]string) {
	for j, fmap := range data {
		fmt.Printf("[%v]\n", j)
		for i, smap := range fmap {
			fmt.Printf("%v=%v\n", i, smap)
		}
		fmt.Println()
	}
}

//PrintIni2 aaa
func PrintIni2(ini map[string]map[string]string) {
	groups := make([]string, 0, len(ini))
	for group := range ini {
		groups = append(groups, group)
	}
	sort.Strings(groups)
	for i, group := range groups {
		fmt.Printf("[%s]\n", group)
		keys := make([]string, 0, len(ini[group]))
		for key := range ini[group] {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Printf("%s=%s\n", key, ini[group][key])
		}
		if i+1 < len(groups) {
			fmt.Println()
		}
	}
}

//MyParsInt a
func MyParsInt(str []string) map[string]map[string]string {
	iniData := map[string]map[string]string{}
	sdata := map[string]string{}
	var key string
	data := false
	for _, l := range str {
		if strings.HasPrefix(l, "[") {
			data = true
			m := strings.TrimFunc(l, func(char rune) bool {
				if char == '[' || char == ']' {
					return true
				}
				return false
			})
			key = m
		}
		if data {
			separ := "="
			if j := strings.Index(l, separ); j > -1 {
				sdata[l[:j]] = l[j+len(separ):]
			} else {
				sdata = map[string]string{}
				iniData[key] = sdata
			}
		}
	}
	return iniData
}

//ParseIni a
func ParseIni(lines []string) map[string]map[string]string {
	const separator = "="
	ini := make(map[string]map[string]string)
	group := "General"
	for _, line := range lines {
		line := strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, ";") {
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			group = line[1 : len(line)-1]
		} else if i := strings.Index(line, separator); i > -1 {
			key := line[:i]
			value := line[i+len(separator):]
			if _, found := ini[group]; !found {
				ini[group] = make(map[string]string)
			}
			ini[group][key] = value
		} else {
			log.Print("failed to parse line:", line)
		}
	}
	return ini
}

//Mat2D aaa
func Mat2D(sl2d []int, col int) [][]int {
	var numStr int
	if len(sl2d)%col != 0 {
		numStr = len(sl2d)/col + 1
	} else {
		numStr = len(sl2d) / col
	}
	matrix := make([][]int, numStr)
	for i := range matrix {
		matrix[i] = make([]int, col)
		copy(matrix[i], sl2d[i*col:])
	}
	return matrix
}

//Smatrix a
func Smatrix(matrix [][]int) []int {
	var Slm []int
	for i := range matrix {
		for _, num := range matrix[i] {
			Slm = append(Slm, num)
		}
	}
	return Slm
}

//UniqueIntSlice UniqueIntSlice
func UniqueIntSlice(array []int) []int {
	var newarray []int
	new := true
	for i := 0; i < len(array); i++ {
		for _, num := range newarray {
			if num == array[i] {
				new = false
				break
			} else {
				new = true
			}
		}
		if new {
			newarray = append(newarray, array[i])
		}
	}
	return newarray
}

//UniqueIntMap UniqueIntMap
func UniqueIntMap(array []int) []int {
	ser := map[int]bool{}
	var newarray []int
	for _, num := range array {
		if _, found := ser[num]; !found {
			newarray = append(newarray, num)
			ser[num] = true
		}
	}
	return newarray
}
