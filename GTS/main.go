package main

import (
	"fmt"
	"os"
	"runtime"

	"./myproj"
)

func main() {
	fmt.Println("GTS started...")

	var str string
	if len(os.Args) == 1 {
		if runtime.GOOS == "windows" {
			str = "D:/md/pti/prSign/"
		} else {
			str = "pr"
		}
	} else {
		str = os.Args[1]
	}

	pr, err := myproj.Xmlpars(str)
	if err != nil {
		fmt.Println("Error in func Xmlpars:" + err.Error())
		return
	}
	err = myproj.MakeAV(pr)
	if err != nil {
		fmt.Println("Error in func MakeAV:" + err.Error())
		return
	}
	fmt.Println("GTS work finished...")
}
