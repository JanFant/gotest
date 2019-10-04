package request

import "sync"

var (
	allstr  = []string{"Baz1:Baz1", "Baz2:Baz2", "DU:DU", "RPU:RPU", "AKNP1:AKNP1", "AKNP2:AKNP2", "AKNP3:AKNP3", "AKNP4:AKNP4"}
	vasdev  = []string{"RPU", "Baz1", "Baz2", "DU"}
	vchsdev = []string{"AKNP1", "AKNP2", "AKNP3", "AKNP4", "RPU"}

	datavas = map[string][]string{"RPU": {"A0IT03IRP", "B0IT03IRP"},
		"Baz1": {"A2IP01IZ1", "B2IP01IZ1", "A0IT01IZ1", "B0IT01IZ1"},
		"Baz2": {"A2IP01IZ2", "B2IP01IZ2", "A0IT02IZ2", "B0IT02IZ2"},
		"DU":   {"A3IP02IDU", "B3IP02IDU"}}
	datavchs = map[string][]string{"AKNP1": {"R0IN01FV1", "R0IN02FV1", "R0IN03FV1"},
		"AKNP2": {"R0IN01FV2", "R0IN02FV2", "R0IN03FV2"},
		"AKNP3": {"R0IN01FV3", "R0IN02FV3", "R0IN03FV3"},
		"AKNP4": {"R0IN01FV4", "R0IN02FV4", "R0IN03FV4"},
		"RPU":   {"R0IN01VRP", "R0IN02VRP", "R0IN03VRP", "R0IN06VRP", "R0IN07VRP"}}

	ipserver     string = "http://192.168.10.30:8080/"
	strAllMB     string = "allModbuses"
	strModval    string = "modvalue?name="
	strModinfo   string = "modinfo?name="
	GlobalData   DevSubLvl1
	modInfofirst bool = false
)

type DevSubLvl1 struct {
	Mute sync.Mutex
	Data []DevSub
}

type DevSub struct {
	Name string
	Sub  []SubValue
}

//SubValue asdasd
type SubValue struct {
	Name  string
	Value []ValueInfo
}

//ValueInfo all info about val
type ValueInfo struct {
	Name string
	Desc string
	val  string
}

//OneData one struct for data
type OneData struct {
	ModbusData   AllModbuses
	ModValueData []AllModvalue
	ModInfoData  []AllModinfo
}

//AllModbuses for allmodbuses
type AllModbuses struct {
	Name     string     `json:"name"`
	Modbuses []Modbuses `json:"Modbuses"`
}

//Modbuses for data modbus info
type Modbuses struct {
	Name    string   `json:"name"`
	Ips     []string `json:"ips"`
	Ports   []int    `json:"ports"`
	Lastops []string `json:"lastop"`
}

//AllModvalue for all modval
type AllModvalue struct {
	Name   string      `json:"name"`
	Values []Valuesmod `json:"values"`
}

//Valuesmod for values
type Valuesmod struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//AllModinfo for data modinfo
type AllModinfo struct {
	Name     string    `json:"name"`
	Registrs []Registr `json:"registers"`
}

//Registr for value mod
type Registr struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
