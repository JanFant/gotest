package myproj

import "encoding/xml"

//Varfile data for varfile
type Varfile struct {
	XMLName xml.Name `xml:"vars"`
	Vars    []Var    `xml:"var"`
}

//Var array variables
type Var struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"description,attr"`
	Format      string `xml:"format,attr"`
}
