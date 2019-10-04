package main

// посмотреть время цикла
// s := time.Now()
//цикл
// e := time.Since(s)
// fmt.Println(e)

import (
	"fmt"

	"./request"
)

func main() {
	request.Firstdata()
	go request.StartQuery()
	// request.Modeselect()
	request.Guimain()
	fmt.Println("end")
}
