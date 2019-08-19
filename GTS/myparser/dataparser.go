package myparser

// Genetal main.xml struct
type Genetal struct {
	Name        string `xml:"name,attr" json:"name"`
	Description string `xml:"description,attr" json:"description"`
	Defdrv      string `xml:"defdrv,attr" json:"defdrv"`
	// Simul       string `xml:"simul,attr" json:"simul"`
	// Ip          string `xml:"ip,attr" json:"ip"`
	// Port        string `xml:"port,attr" json:"port"`
	Subs []Subs `xml:"subs" json:"subs"`
}

// Subs main.xml struct
type Subs struct {
	Name string `xml:"name,attr" json:"name"`
	// Scheme      string `xml:"scheme,attr" json:"scheme"`
	Id   string `xml:"id,attr" json:"id"`
	File string `xml:"file,attr" json:"file"`
	// Description string `xml:"description,attr" json:"decription"`
	// Main        string `xml:"main,attr" json:"main"`
	// Second      string `xml:"second,attr" json:"second"`
	Path string `xml:"path,attr" json:"path"`
	// Step        string `xml:"step,attr" json:"step"`
}

// DrTable  all driver includ defdrv
type DrTable struct {
	Name string  `xml:"name,attr"`
	Bsig Signals `xml:"signals"`
}

//Signals signals in drv
type Signals struct {
	Msig []Signal `xml:"signal"`
}

//Signal signal in driver
type Signal struct {
	Chan   string `xml:"name,attr"`
	Format string `xml:"format,attr"`
}
