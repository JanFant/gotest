package myparser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

//Xmlpars mainparser
func Xmlpars(path string) {
	var DataMain Genetal                // parsed Main.xml
	DataDrv := make(map[string]DrTable) // parsed defdrv

	namefile := path + "main.xml"
	datafile, err := ioutil.ReadFile(namefile)
	if err != nil {
		fmt.Println("Error read file main.xml: " + err.Error())
		return
	}
	err = xml.Unmarshal(datafile, &DataMain)
	if err != nil {
		fmt.Println("Error unmarshal main.xml: " + err.Error())
		return
	}

	//parser for defdrv
	DrTablePath := path + DataMain.Defdrv // path to all driver tables
	DataDrv, err = DataTbpars(DrTablePath)
	if err != nil {
		fmt.Println("Error parser devdrv :" + err.Error())
	}

	println(DataDrv["fds16"].Name)

	fmt.Println("fin")
}

//DataTbpars data driver parser
func DataTbpars(DrtabPath string) (map[string]DrTable, error) {
	tempDataDrv := make(map[string]DrTable) //Create temp map to all driver tables
	dataTfile, err := ioutil.ReadDir(DrtabPath)
	if err != nil {
		fmt.Println("Error Driver Dir: " + err.Error())
		return nil, err
	}

	for i := range dataTfile {
		var TempData DrTable
		nameTfile := DrtabPath + "/" + dataTfile[i].Name()

		readfile, err := ioutil.ReadFile(nameTfile)
		if err != nil {
			fmt.Println("Error read drv file " + nameTfile + ":" + err.Error())
			return nil, err
		}

		err = xml.Unmarshal(readfile, &TempData)
		if err != nil {
			fmt.Println("Error unmarshal file" + nameTfile + ":" + err.Error())
			return nil, err
		}

		tempDataDrv[TempData.Name] = TempData
	}
	return tempDataDrv, err
}
