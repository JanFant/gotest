package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
)

// Genetal asd
type Genetal struct {
	Name        string `xml:"name,attr" json:"name"`
	Description string `xml:"description,attr" json:"description"`
	Defdrv      string `xml:"defdrv,attr" json:"defdrv"`
	Simul       string `xml:"simul,attr" json:"simul"`
	Ip          string `xml:"ip,attr" json:"ip"`
	Port        string `xml:"port,attr" json:"port"`
	Subs        []Subs `xml:"subs" json:"subs"`
}

// Subs asd
type Subs struct {
	Name        string `xml:"name,attr" json:"name"`
	Scheme      string `xml:"scheme,attr" json:"scheme"`
	Id          string `xml:"id,attr" json:"id"`
	File        string `xml:"file,attr" json:"file"`
	Description string `xml:"description,attr" json:"decription"`
	Main        string `xml:"main,attr" json:"main"`
	Second      string `xml:"second,attr" json:"second"`
	Path        string `xml:"path,attr" json:"path"`
	Step        string `xml:"step,attr" json:"step"`
}

func main() {
	nfile := "main"
	filexml, err := ioutil.ReadFile(nfile + ".xml")
	if err != nil {
		println(err)
	}
	var datatest Genetal

	err = xml.Unmarshal(filexml, &datatest)
	if err != nil {
		println(err)
	}

	filejson, err := json.Marshal(datatest)
	if err != nil {
		println(err)
	}

	err = ioutil.WriteFile(nfile+".json", filejson, 0644)
	if err != nil {
		println(err)
	}
	println("Finish!")
}
