package myproj

// Genetal main.xml struct
type Genetal struct {
	Name        string `xml:"name,attr" json:"name"`
	Description string `xml:"description,attr" json:"description"`
	Defdrv      string `xml:"defdrv,attr" json:"defdrv"`
	// Simul       string `xml:"simul,attr" json:"simul"`
	// Ip          string `xml:"ip,attr" json:"ip"`
	// Port        string `xml:"port,attr" json:"port"`
	Subs  []Subs `xml:"subs" json:"subs"`
	DrTab map[string]DrTable
	Path  string
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
	Data *FpDev
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

//FpDev asd
type FpDev struct {
	Fp  DataFp  `xml:"subsystem"`
	Dev DataDev `xml:"devices"`
}

//DataFp data for *fp.xml
type DataFp struct {
	Variable Variable `xml:"variable"`
	Devices  Devices  `xml:"devices"`
}

//Variable Variable
type Variable struct {
	Name string `xml:"xml,attr"`
}

//Devices Devices
type Devices struct {
	Name string `xml:"xml,attr"`
}

//DataDev data for dev*.xml
type DataDev struct {
	Name string   `xml:"xml,attr"`
	Dev  []Device `xml:"device"`
}

//Device Device
type Device struct {
	Name  string `xml:"name,attr"`
	Drv   string `xml:"driver,attr"`
	Slot  string `xml:"slot,attr"`
	Descr string `xml:"description,attr"`
}
