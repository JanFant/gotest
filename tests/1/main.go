package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := Minimum(4, 3, 8, 2, 9).(int)
	fmt.Printf("%T %v\n", i, i)
	f := Minimum(9.4, -5.4, 3.8, 17.0, -3.1, 0.0).(float64)
	fmt.Printf("%T %v\n", f, f)
	s := Minimum("K", "X", "B", "C", "CC", "CA", "D", "M").(string)
	fmt.Printf("%T %q\n", s, s)
	fmt.Println("===================================")
	xs := []int{2, 4, 6, 8}
	fmt.Println("5 @", Index(xs, 5), " 6 @", Index(xs, 6))
	ys := []string{"C", "B", "K", "A"}
	fmt.Println("Z @", Index(ys, "Z"), " A @", Index(ys, "A"))

	fmt.Println("===================================")
	fmt.Println("5 @", IndexReflectX(xs, 5), " 6 @", IndexReflectX(xs, 6))
	fmt.Println("Z @", IndexReflectX(ys, "Z"), " A @", IndexReflectX(ys, "A"))

	fmt.Println("===================================")
	fmt.Println("5 @", IndexReflect(xs, 5), " 6 @", IndexReflect(xs, 6))
	fmt.Println("Z @", IndexReflect(ys, "Z"), " A @", IndexReflect(ys, "A"))

}

func Minimum(first interface{}, rest ...interface{}) interface{} {
	minimum := first
	for _, x := range rest {
		switch x := x.(type) {
		case int:
			if x < minimum.(int) {
				minimum = x
			}
		case float64:
			if x < minimum.(float64) {
				minimum = x
			}
		case string:
			if x < minimum.(string) {
				minimum = x
			}
		}
	}
	return minimum
}

func Index(xs interface{}, x interface{}) int {
	switch slice := xs.(type) {
	case []int:
		for i, y := range slice {
			if y == x.(int) {
				return i
			}
		}
	case []string:
		for i, y := range slice {
			if y == x.(string) {
				return i
			}
		}
	}
	return -1
}

func IndexReflectX(xs interface{}, x interface{}) int { // Более длинное решение
	if slice := reflect.ValueOf(xs); slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++ {
			switch y := slice.Index(i).Interface().(type) {
			case int:
				if y == x.(int) {
					return i
				}
			case string:
				if y == x.(string) {
					return i
				}
			}
		}
	}
	return -1
}

func IndexReflect(xs interface{}, x interface{}) int {
	if slice := reflect.ValueOf(xs); slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++ {
			if reflect.DeepEqual(x, slice.Index(i)) {
				return i
			}
		}
	}
	return -1
}
