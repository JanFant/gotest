package request

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

//StartQuery start
func StartQuery() {
	for {
		data := QueryData()
		parserData(data)
		time.Sleep(time.Second / 2)
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
		if !modInfofirst {
			body = MakeRequest(ipserver + strModinfo + str)
			json.Unmarshal(body, &data.ModInfoData[num])
		}
	}
	return data
}

func parserData(moddata OneData) {
	var numSub int
	// GlobalData.Mute.Lock()
	// defer GlobalData.Mute.Unlock()
	for num1, dev := range GlobalData.Data {
		for num2, sub := range dev.Sub {
			//ищем какой номер девайса
			for numstr, str := range allstr {
				if strings.Contains(str, sub.Name) {
					numSub = numstr
					break
				}
			}
			//-----
			for num3, value := range sub.Value {
				//------
				for _, valueMod := range moddata.ModValueData[numSub].Values {
					if value.Name == valueMod.Name {
						GlobalData.Mute.Lock()
						GlobalData.Data[num1].Sub[num2].Value[num3].val = valueMod.Value
						GlobalData.Mute.Unlock()
						break
					}
				}
				if !modInfofirst {
					for _, valueInf := range moddata.ModInfoData[numSub].Registrs {
						if value.Name == valueInf.Name {
							GlobalData.Mute.Lock()
							GlobalData.Data[num1].Sub[num2].Value[num3].Desc = valueInf.Desc
							GlobalData.Mute.Unlock()
							break
						}
					}
					modInfofirst = true
				}
				//-----
			}
		}
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

func Firstdata() {
	GlobalData.Data = make([]DevSub, 2)
	for i := 0; i < 2; i++ {
		switch {
		case i == 0:
			makedata(&GlobalData.Data[i], "vas", vasdev, datavas)
		case i == 1:
			makedata(&GlobalData.Data[i], "vchs", vchsdev, datavchs)
		}
	}

}

func makedata(data *DevSub, name string, devise []string, maphead map[string][]string) {
	data.Sub = make([]SubValue, len(devise))
	data.Name = name
	for num, dev := range devise {
		data.Sub[num].Name = dev
		for numm, val := range maphead[dev] {
			if numm == 0 {
				data.Sub[num].Value = make([]ValueInfo, len(maphead[dev]))
			}
			data.Sub[num].Value[numm].Name = val
		}
	}
}
