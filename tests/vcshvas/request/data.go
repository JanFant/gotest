package request

var (
	allstr = []string{"Baz1:Baz1", "Baz2:Baz2", "DU:DU", "RPU:RPU", "AKNP1:AKNP1", "AKNP2:AKNP2", "AKNP3:AKNP3", "AKNP4:AKNP4"}
	data   = map[string][]string{"rpu": {"A0IT03IRP", "B0IT03IRP"},
		"baz1": {"A2IP01IZ1", "B2IP01IZ1", "A0IT01IZ1", "B0IT01IZ1"},
		"baz2": {"A2IP01IZ2", "B2IP01IZ2", "A0IT01IZ2", "B0IT01IZ2"},
		"du":   {"A3IP02IDU", "B3IP02IDU"}}

	aknp1Vchs = map[string][]string{"R0IN01FV1", "R0IN02FV1", "R0IN03FV1"}
	aknp2Vchs = map[string][]string{"R0IN01FV2", "R0IN02FV2", "R0IN03FV2"}
	aknp3Vchs = map[string][]string{"R0IN01FV3", "R0IN02FV3", "R0IN03FV3"}
	aknp4Vchs = map[string][]string{"R0IN01FV4", "R0IN02FV4", "R0IN03FV4"}
	rpuVchs   = map[string][]string{"R0IN01VRP", "R0IN02VRP", "R0IN03VRP", "R0IN06VRP", "R0IN07VRP"}

	ipserver   string = "http://192.168.10.30:8080/"
	strAllMB   string = "allModbuses"
	strModval  string = "modvalue?name="
	strModinfo string = "modinfo?name="
)

//DeviceInfo divice vas vchs
type DeviceInfo struct {
	Sub map[string]map[string]ValueInfo
}

//ValueInfo all info about val
type ValueInfo struct {
	Desc string
	val  int
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
