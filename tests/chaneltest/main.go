package main

import (
	"math/rand"
)

func main() {
	ch1 := make(chan int)
	go funCh1(ch1)
}

func funCh1(cn chan<- int) {
	for {
		cn <- rand.Int()
	}
}

func funCh2() chan int {
	for {
	}
}
