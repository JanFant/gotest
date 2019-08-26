package main

import (
	"math"
	"sort"
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
	stdDev  float64
	mode    []float64
}

func mode(stat *statistics) {
	var (
		temp []float64
	)
	a := make(map[float64]int)
	k := 0
	sort.Float64s(stat.numbers)
	i := 0
	for i = range stat.numbers {
		if i == (len(stat.numbers) - 1) {
			temp = append(temp, stat.numbers[i])
			temp = append(temp, float64(k+1))
			a[stat.numbers[i]] = k + 1
			break
		}
		if stat.numbers[i] == stat.numbers[i+1] {
			k++
		} else {
			temp = append(temp, stat.numbers[i-1])
			temp = append(temp, float64(k+1))
			a[stat.numbers[i-1]] = k + 1
			k = 0
		}
	}
	for i < len(a) {
		math.Max(
	}
	// o := sort.SearchFloat64s(stat.numbers, stat.numbers[1])
	// println(o)
	// for range stat.numbers {
	// 	if m >= len(stat.numbers) {
	// 		break
	// 	}
	// 	freq[k*2] = stat.numbers[m]
	// 	freq[k*2+1]++
	// 	m++
	// 	for m < len(stat.numbers) {
	// 		if freq[k*2] == stat.numbers[m] {
	// 			freq[k*2+1]++
	// 			m++
	// 		} else {
	// 			break
	// 		}
	// 	}
	// 	k++
	// }

	// m = 0

}

func main() {
	var a statistics
	a.numbers = append(a.numbers, 1, 2, 1, 1, 2, 2, 1, 3, 3, 3, 3, 5)

	mode(&a)
}
