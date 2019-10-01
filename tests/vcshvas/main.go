package main

// посмотреть время цикла
// s := time.Now()
//цикл
// e := time.Since(s)
// fmt.Println(e)

import (
	"fmt"
	"time"

	"./netsite"
	"./request"
)

func main() {
	go netsite.NetVchsVas()
	for {
		data := request.QueryData()
		time.Sleep(time.Second / 2)
		fmt.Println(data.ModbusData.Modbuses[0].Lastops[0] + "    " + time.Now().String())
	}
	fmt.Println("end")
}
