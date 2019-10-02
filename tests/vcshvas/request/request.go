package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//StartQuery start
func StartQuery() {
	for {
		data := QueryData()
		time.Sleep(time.Second / 2)
		fmt.Println(data.ModbusData.Modbuses[0].Lastops[0] + "    " + time.Now().String())
	}
}

//QueryData query json and pars it
func QueryData() OneData {
	var data OneData
	data.ModValueData = make([]AllModvalue, len(allstr))
	data.ModInfoData = make([]AllModinfo, len(allstr))
	body := MakeRequest(ipserver + strAllMB)
	json.Unmarshal(body, &data.ModbusData)
	for num, str := range allstr {
		body := MakeRequest(ipserver + strModval + str)
		json.Unmarshal(body, &data.ModValueData[num])
		body = MakeRequest(ipserver + strModinfo + str)
		json.Unmarshal(body, &data.ModInfoData[num])
	}
	return data
}

func parserData(data OneData) {
	var ParsData [2]DeviceInfo
	for num, str := range allstr {
		
	}
}

//MakeRequest query json from server
func MakeRequest(str string) []byte {
	resp, err := http.Get(str)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}
