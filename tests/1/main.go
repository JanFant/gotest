package main

import (
	"fmt"
	"math"
)

func main() {
	xs := []int{2, 4, 6, 8}
	ys := []string{"C", "B", "K", "A"}
	fmt.Println(
		SliceIndex(len(xs), func(i int) bool { return xs[i] == 5 }),
		SliceIndex(len(xs), func(i int) bool { return xs[i] == 6 }),
		SliceIndex(len(ys), func(i int) bool { return ys[i] == "Z" }),
		SliceIndex(len(ys), func(i int) bool { return ys[i] == "A" }))

	i := SliceIndex(math.MaxInt32,
		func(i int) bool { return i > 0 && i%27 == 0 && i%51 == 0 })
	fmt.Println(i)
	fmt.Println("------------------------------------")
	readings := []int{4, -3, 2, -7, 8, 19, -11, 7, 18, -6}
	even := IntFilter(readings, func(i int) bool { return i < 0 && i%2 == 0 })
	fmt.Println(even)
	fmt.Println("------------------------------------")

	readings = []int{4, -3, 2, -7, 8, 19, -11, 7, 18, -6}
	even = make([]int, 0, len(readings))
	Filter(len(readings), func(i int) bool { return readings[i]%2 == 0 },
		func(i int) { even = append(even, readings[i]) })
	fmt.Println(even)

}

//SliceIndex aa
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

//IntFilter  aa
func IntFilter(slice []int, predicate func(int) bool) []int {
	filtered := make([]int, 0, len(slice))
	for i := 0; i < len(slice); i++ {
		if predicate(slice[i]) {
			filtered = append(filtered, slice[i])
		}
	}
	return filtered
}

//Filter aaa
func Filter(limit int, predicate func(int) bool, appender func(int)) {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			appender(i)
		}
	}
}
