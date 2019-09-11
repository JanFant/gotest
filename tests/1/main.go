package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

//293
func main() {
	// for i := 1; i <= 4; i++ {
	// 	a, b, c := PythagoreanTriple(i, i+1)
	// 	q1 := Heron(a, b, c)
	// 	q2 := Heron(PythagoreanTriple(i, i+1))
	// 	fmt.Printf("q1 == %10f == q2 == %10f\n", q1, q2)
	// }
	//palindrom
	str := "А роза упала на лапу Азора"
	fmt.Println("Word =", str, " is palindrom =", Palindrom(str))
	//hofstadter
	// females := make([]int, 20)
	// males := make([]int, len(females))
	// for n := range females {
	// 	females[n] = HofstadterFemale(n)
	// 	males[n] = HofstadterMale(n)
	// }
	// fmt.Println("F", females)
	// fmt.Println("M", males)

	//fibo
	// for n := 0; n < 20; n++ {
	// 	fmt.Print(Fibonacci(n), " ")
	// }
	// fmt.Println()
}

func Palindrom(str string) bool {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ToLower(str)
	return Pali(str)
}

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

func HofstadterFemale(n int) int {
	if n <= 0 {
		return 1
	}
	return n - HofstadterMale(HofstadterFemale(n-1))
}

func HofstadterMale(n int) int {
	if n <= 0 {
		return 0
	}
	return n - HofstadterFemale(HofstadterMale(n-1))
}

func Heron(a, b, c int) float64 {
	a1, b1, c1 := float64(a), float64(b), float64(c)
	s := (a1 + b1 + c1) / 2
	return math.Sqrt(s * (s - a1) * (s - b1) * (s - c1))
}

func PythagoreanTriple(m, n int) (a, b, c int) {
	if m < n {
		m, n = n, m
	}
	return (m * m) - (n * n), (2 * m * n), (m * m) + (n * n)
}

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
