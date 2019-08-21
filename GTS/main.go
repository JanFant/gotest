package main

import (
	"fmt"

	"./myparser"
)

const str = "D:/space/1/pr/"

// const str = "E:/workfolder/Pr/"

func main() {
	pr, err := myparser.Xmlpars(str)
	if err != nil {
		fmt.Println("Error in func Xmlpars:" + err.Error())
	}
	err = myparser.MakeAV(pr)
	if err != nil {
		fmt.Println("Error in func MakeAV:" + err.Error())
	}
	// fmt.Println(pr.Name)
}
