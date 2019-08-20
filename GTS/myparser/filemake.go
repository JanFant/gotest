package myparser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

//MakeAV asd
func MakeAV(Data *Genetal) error {

	err := VarMake(Data)
	if err != nil {
		fmt.Println("Error VarMake :" + err.Error())
		return err
	}

	// VarMake(Data, 0)
	return err
}

//VarMake make var.xml files
func VarMake(Data *Genetal) error {
	var err error
	for _, sub := range Data.Subs {
		var (
			tempVFile Varfile
			tempVar   Var
			varname   string
		)
		for _, dev := range sub.Data.Dev.Dev {
			for _, sig := range Data.DrTab[dev.Drv].Bsig.Msig {
				tempVar.Name = dev.Name + sig.Chan
				tempVar.Description = "(" + dev.Descr + " - " + sig.Chan + ")"
				tempVar.Format = sig.Format
				tempVFile.Vars = append(tempVFile.Vars, tempVar)
			}
		}
		varname = "var/" + sub.Data.Fp.Variable.Name + ".xml"
		err = SaveMarshalfile(varname, tempVFile)
		if err != nil {
			fmt.Println("Error SaveMarshalfile :" + varname + " " + err.Error())
			return err
		}
	}
	return err
}

//SaveMarshalfile marshal and save file :)
func SaveMarshalfile(namefile string, file interface{}) error {
	tempfile, err := xml.Marshal(file)
	if err != nil {
		fmt.Println("Error Marshal file :" + namefile + " " + err.Error())
		return err
	}

	err = ioutil.WriteFile(namefile, tempfile, 0644)
	if err != nil {
		fmt.Println("Error WriteFile :" + namefile + " " + err.Error())
		return err
	}
	return err
}

//AssMake ...
func AssMake() {

}
