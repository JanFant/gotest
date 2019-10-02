package main

// посмотреть время цикла
// s := time.Now()
//цикл
// e := time.Since(s)
// fmt.Println(e)

import (
	"fmt"

	"./netsite"
	"./request"
)

func main() {
	go request.StartQuery()
	netsite.NetVchsVas()

	fmt.Println("end")
}
