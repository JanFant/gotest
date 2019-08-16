package main

import (
	"encoding/json"
	"io/ioutil"
)

//Updata ayaya
type Updata struct {
	Ndata []Data `json:"rows"`
}

//Data struct for test
type Data struct {
	Id                 string  `json:"id"`
	AccountId          string  `json:"accountId"`
	MinimumBalance     float32 `json:"minimumBalance"`
	ModificationsCount int64   `json:"modificationsCount"`
	IsSerialTrackable  bool    `json:"isSerialTrackable"`
	Stock              float32 `json:"stock"`
	Reserve            float32 `json:"reserve"`
	InTransit          float32 `json:"inTransit"`
	Quantity           float32 `json:"quantity"`
}

func main() {

	file, err := ioutil.ReadFile("test.json")
	if err != nil {
		println(err)
	}

	examp := new(Updata)
	if err := json.Unmarshal(file, &examp); err != nil {
		println(err)
	}

	i := 1
	examp.Ndata[i].Id = "0"
	examp.Ndata[i].AccountId = "0"
	examp.Ndata[i].MinimumBalance = 322.228
	examp.Ndata[i].ModificationsCount = 1
	examp.Ndata[i].IsSerialTrackable = false
	examp.Ndata[i].Stock = 1.01
	examp.Ndata[i].Reserve = 1.01
	examp.Ndata[i].InTransit = 2.3
	examp.Ndata[i].Quantity = 32

	fin, err := json.Marshal(examp)
	if err != nil {
		println(err)
	}
	ioutil.WriteFile("test1.json", fin, 0644)
	println("AYAYA")
}
