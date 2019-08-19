package myparser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

//Xmlpars mainparser
func Xmlpars(path string) {
	namefile := path + "main.xml"
	datafile, err := ioutil.ReadFile(namefile)
	if err != nil {
		fmt.Println("Error read file main.xml: " + err.Error())
		return
	}
	var DataMain Genetal
	err = xml.Unmarshal(datafile, &DataMain)
	if err != nil {
		fmt.Println("Error unmarshal main.xml: " + err.Error())
		return
	}

	// read all xml file in defdrv
	DrTablePath := path + DataMain.Defdrv
	dataTfile, err := ioutil.ReadDir(DrTablePath)
	if err != nil {
		fmt.Println("Error Driver Dir: " + err.Error())
		return
	}

	//Create map for drive table
	DataDrv := make(map[string]DrTable)
	for i := range dataTfile {
		nameTfile := DrTablePath + "/" + dataTfile[i].Name()
		tempDrv, err := DataTablepars(nameTfile)
		if err != nil {
			fmt.Println("Error DrvData: " + err.Error())
			return
		}
		DataDrv[tempDrv.Name] = *tempDrv
	}
	fmt.Println("fin")
	fmt.Println("asd")
}

//DataTablepars data driver parser (*DrTable for correct return)
func DataTablepars(NameFile string) (*DrTable, error) {
	readfile, err := ioutil.ReadFile(NameFile)
	if err != nil {
		fmt.Println("Error read drv file " + NameFile + ":" + err.Error())
		return nil, err
	}
	var TempData DrTable
	err = xml.Unmarshal(readfile, &TempData)
	if err != nil {
		fmt.Println("Error unmarshal file" + NameFile + ":" + err.Error())
		return nil, err
	}
	return &TempData, err
}
