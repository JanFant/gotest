package main

import (
	"fmt"
	"sort"
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
	stdDev  float64
	mode    []float64
}

func mode(stat statistics) []float64 {
	var (
		temp []float64
	)
	k := 1
	pastK := 0
	pastnNum := 0.0
	sort.Float64s(stat.numbers)
	for i, num := range stat.numbers {
		if i == 0 {
			pastnNum = num
			k = 1
			continue
		} else {
			if pastnNum == num {
				k++
			} else {
				if pastK > k {
					k = 1
					pastnNum = num
					continue
				} else if pastK == k {
					temp = append(temp, pastnNum)
					k = 1
				} else {
					temp = temp[:0]
					temp = append(temp, pastnNum)
					pastK = k
					k = 1
				}
			}
			pastnNum = num
		}
	}
	if pastK == k {
		temp = append(temp, pastnNum)
	} else if k > pastK {
		temp = temp[:0]
		temp = append(temp, pastnNum)
	}

	return temp
}

func main() {
	var a statistics
	a.numbers = append(a.numbers, 1, 2, 1, 1, 1, 2, 2, 1, 3, 2, 3, 3, 3, 3, 5, 5, 5, 5)

	fmt.Println(mode(a))
}
