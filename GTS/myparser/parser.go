package myparser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

//Xmlpars mainparser
func Xmlpars(path string) {
	var DataMain *Genetal               // parsed Main.xml
	DataDrv := make(map[string]DrTable) // parsed defdrv

	namefile := path + "main.xml"
	DataMain, err := PMainXML(namefile)
	if err != nil {
		fmt.Println("Error parser MainProj :" + err.Error())
	}

	//parser for subs
	for _, subs := range DataMain.Subs {
		FpPath := path + subs.Path + "/" + subs.File + ".xml"
		PSubsXML(FpPath)
	}

	//parser for defdrv
	DrTablePath := path + DataMain.Defdrv // path to all driver tables
	DataDrv, err = PDefdrvXML(DrTablePath)
	if err != nil {
		fmt.Println("Error parser devdrv :" + err.Error())
		return
	}

	println(DataDrv["fds16"].Name)

	fmt.Println("fin")
}

//PSubsXML parser *fp.xml + dev*.xml
func PSubsXML(FpPath string) (*Genetal, error) {
	var tempData FpDev
	datafile, err := ioutil.ReadFile(FpPath)
	if err != nil {
		fmt.Println("Error read file :" + FpPath + " - " + err.Error())
		return nil, err
	}
	err = xml.Unmarshal(datafile, &tempData.Fp)
	if err != nil {
		fmt.Println("Error unmarshal : " + FpPath + " - " + err.Error())
		return nil, err
	}

	return nil, err
}

//PMainXML parser Main.xml
func PMainXML(mainpath string) (*Genetal, error) {
	var tempData Genetal
	datafile, err := ioutil.ReadFile(mainpath)
	if err != nil {
		fmt.Println("Error read file main.xml: " + err.Error())
		return nil, err
	}
	err = xml.Unmarshal(datafile, &tempData)
	if err != nil {
		fmt.Println("Error unmarshal main.xml: " + err.Error())
		return nil, err
	}
	return &tempData, err
}

//PDefdrvXML data driver parser
func PDefdrvXML(DrtabPath string) (map[string]DrTable, error) {
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
