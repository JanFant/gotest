package request

var (
	allstr           = []string{"Baz1:Baz1", "Baz2:Baz2", "DU:DU", "RPU:RPU", "AKNP1:AKNP1", "AKNP2:AKNP2", "AKNP3:AKNP3", "AKNP4:AKNP4"}
	ipserver  string = "http://192.168.10.30:8080/"
	strAllMB  string = "allModbuses"
	strModval string = "modvalue?name="
)

//OneData one struct for data
type OneData struct {
	ModbusData   AllModbuses
	ModValueData []AllModvalue
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
