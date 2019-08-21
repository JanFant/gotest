package myproj

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

//Xmlpars mainparser
func Xmlpars(path string) (*Genetal, error) {
	var DataMain *Genetal // parsed Main.xml
	namefile := path + "main.xml"
	DataMain, err := PMainXML(namefile)
	if err != nil {
		fmt.Println("Error parser :" + namefile + " - " + err.Error())
		return nil, err
	}
	DataMain.Path = path
	//parser for subs
	for i, subs := range DataMain.Subs {
		FpPath := path + subs.Path + "/"
		DataMain.Subs[i].Data, err = PSubsXML(FpPath, subs.File)
		if err != nil {
			fmt.Println("Error parser :" + FpPath + " - " + err.Error())
			return nil, err
		}
	}

	//parser for defdrv
	DrTablePath := path + DataMain.Defdrv // path to all driver tables
	DataMain.DrTab, err = PDefdrvXML(DrTablePath)
	if err != nil {
		fmt.Println("Error parser devdrv :" + err.Error())
		return nil, err
	}
	return DataMain, err
}

//PSubsXML parser *fp.xml + dev*.xml
func PSubsXML(FpPathm, FpName string) (*FpDev, error) {
	var tempData FpDev
	// fp.xml
	str := FpPathm + FpName + ".xml"
	datafile, err := ioutil.ReadFile(str)
	if err != nil {
		fmt.Println("Error read file :" + str + " - " + err.Error())
		return nil, err
	}
	err = xml.Unmarshal(datafile, &tempData.Fp)
	if err != nil {
		fmt.Println("Error unmarshal : " + str + " - " + err.Error())
		return nil, err
	}
	// dev.xml
	str = FpPathm + tempData.Fp.Devices.Name + ".xml"
	datafile, err = ioutil.ReadFile(str)
	if err != nil {
		fmt.Println("Error read file :" + str + " - " + err.Error())
		return nil, err
	}
	err = xml.Unmarshal(datafile, &tempData.Dev)
	if err != nil {
		fmt.Println("Error unmarshal : " + str + " - " + err.Error())
		return nil, err
	}

	return &tempData, err
}

//PMainXML parser Main.xml
func PMainXML(mainpath string) (*Genetal, error) {
	var tempData Genetal
	datafile, err := ioutil.ReadFile(mainpath)
	if err != nil {
		fmt.Println("Error read file :" + mainpath + " - " + err.Error())
		return nil, err
	}
	err = xml.Unmarshal(datafile, &tempData)
	if err != nil {
		fmt.Println("Error unmarshal :" + mainpath + " - " + err.Error())
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
