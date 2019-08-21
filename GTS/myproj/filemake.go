package myproj

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/beevik/etree"
	"github.com/clbanning/mxj"
)

//MakeAV make var and ass file
func MakeAV(Data *Genetal) error {
	err := VarMake(Data)
	if err != nil {
		fmt.Println("Error VarMake :" + err.Error())
		return err
	}
	err = AssMake(Data)
	if err != nil {
		fmt.Println("Error AssMake :" + err.Error())
		return err
	}
	return err
}

//VarMake make *var.xml files
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
		varname = Data.Path + sub.Name + "/" + sub.Data.Fp.Variable.Name + ".xml"
		// netblock for goteck
		tempVar.Name = "netblock"
		tempVar.Description = "( netblock )"
		tempVar.Format = "1"
		tempVFile.Vars = append(tempVFile.Vars, tempVar)

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
	tempfile, err = mxj.BeautifyXml(tempfile, "", "\t")
	if err != nil {
		fmt.Println("Error BeautifyXml file :" + namefile + " " + err.Error())
		return err
	}
	err = ioutil.WriteFile(namefile, tempfile, 0644)
	if err != nil {
		fmt.Println("Error WriteFile :" + namefile + " " + err.Error())
		return err
	}
	return err
}

//AssMake make *ass.xml file
func AssMake(Data *Genetal) error {
	var err error
	for _, sub := range Data.Subs {
		file := etree.NewDocument()
		assign := file.CreateElement("assign")
		for _, dev := range sub.Data.Dev.Dev {
			devname := assign.CreateElement(dev.Name)
			for _, sig := range Data.DrTab[dev.Drv].Bsig.Msig {
				def := devname.CreateElement("def")
				def.CreateAttr("name", dev.Name+sig.Chan)
				def.CreateText(sig.Chan)
			}
		}
		file.Indent(2)
		namefile := Data.Path + sub.Name + "/" + sub.Data.Dev.Name + ".xml"
		err = file.WriteToFile(namefile)
		if err != nil {
			fmt.Println("Error WriteToFile :" + namefile + " - " + err.Error())
			return err
		}
	}
	return err
}
