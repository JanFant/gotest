package request

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//QueryData query json and pars it
func QueryData() OneData {
	var data OneData
	data.ModValueData = make([]AllModvalue, len(allstr))
	body := MakeRequest(ipserver + strAllMB)
	json.Unmarshal(body, &data.ModbusData)
	for num, str := range allstr {
		body1 := MakeRequest(ipserver + strModval + str)
		json.Unmarshal(body1, &data.ModValueData[num])
	}
	return data
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
