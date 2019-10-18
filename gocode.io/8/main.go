package main

import (
	"fmt"
	"reflect"
)

func main() {
	println("Logging in...")
	authorized := startup(login())
	if reflect.ValueOf(authorized).Bool() {
		println("Starting the engine")
		return
	}
	println("Startup failed")
}

func validSequence(i int, el interface{}) bool {
	fmt.Println(reflect.ValueOf(el).Elem().Field(0).Interface())

	return reflect.TypeOf(el).String() == "*main.Sequence" &&
		!reflect.ValueOf(el).IsNil() &&
		reflect.ValueOf(el).Elem().NumField() == 2 &&
		reflect.TypeOf(reflect.ValueOf(el).Elem().Field(0).Interface()).String() == "int" &&
		int(reflect.ValueOf(el).Elem().Field(0).Int()) == i*i-i &&
		!reflect.ValueOf(reflect.ValueOf(el).Elem().Field(1).Interface()).IsNil()
}

func startup(seq interface{}) bool {
	for i := 0; i < 5; i++ {
		if !validSequence(i, seq) {
			return false
		}
		seq = reflect.ValueOf(seq).Elem().Field(1).Interface()
	}

	return true
}

type Sequence struct {
	A1 int
	A2 []int
}

func login() *Sequence {
	var a Sequence
	a.A1 = 0
	for i := 0; i < 5; i++ {
		a.A2 = append(a.A2, int(i*i-i))
	}
	return &a
}
