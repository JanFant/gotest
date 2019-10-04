package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//StartQuery start
func StartQuery() {
	for {
		data := QueryData()
		parserData(data)
		time.Sleep(time.Second / 2)
		// fmt.Println(data.ModbusData.Modbuses[0].Lastops[0] + "    " + time.Now().String())
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
						// GlobalData.Data[num1].Sub[num2].Value[num3].val, _ = strconv.Atoi(strings.TrimSpace(valueMod.Value))
						GlobalData.Data[num1].Sub[num2].Value[num3].val = valueMod.Value
						// GlobalData.Data[num1].Sub[num2].Value[num3].val, _ = strconv.ParseFloat(strings.TrimSpace(valueMod.Value), 64)
						GlobalData.Mute.Unlock()
						break
					}
				}
				for _, valueInf := range moddata.ModInfoData[numSub].Registrs {
					if value.Name == valueInf.Name {
						GlobalData.Mute.Lock()
						GlobalData.Data[num1].Sub[num2].Value[num3].Desc = valueInf.Desc
						GlobalData.Mute.Unlock()
						break
					}
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

func Modeselect() {
	var dev string
	fmt.Println("Для выхода из программы quit")
	for {
		fmt.Print("Введите какое устройство (vchs, vas):")
		fmt.Scan(&dev)
		switch {
		case dev == "vas":
			GlobalData.Mute.Lock()
			devwork(GlobalData.Data[0], vasdev, 0)
			GlobalData.Mute.Unlock()
			fmt.Println(dev)
		case dev == "vchs":
			GlobalData.Mute.Lock()
			devwork(GlobalData.Data[1], vchsdev, 1)
			GlobalData.Mute.Unlock()
		case dev == "quit":
			return
		default:
			fmt.Println("не верно введено устройство")

		}
	}
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

func devwork(data DevSub, dev []string, numdevglobal int) {
	var (
		strdev string
		sub    string
	)
	for num, dev := range dev {
		strdev = strdev + fmt.Sprint(num, " - ", dev, ";  ")
	}
	for {
		fmt.Print("введите номер подсистемы : ", strdev, ": ")
		fmt.Scan(&sub)
		numdev, _ := strconv.Atoi(sub)
		if numdev > len(dev) {
			fmt.Println("указан неверный номер устройства")
			continue
		} else if sub == "quit" {
			break
		} else {
			var strval []string
			for numval, val := range data.Sub[numdev].Value {
				tempstr := fmt.Sprint(numval, " - ", val.Name, " ", val.Desc)
				strval = append(strval, tempstr)
			}
			fmt.Println("введите номер переменной : ")
			for _, str := range strval {
				fmt.Println(str)
			}
			fmt.Scan(&sub)
			numval, _ := strconv.Atoi(sub)
			if numval > len(vasdev) {
				fmt.Println("указан неверный номер устройства")
				continue
			} else if sub == "quit" {
				break
			} else {
				table(numdevglobal, numdev, numval)
			}
		}
	}
}

func table(num1, num2, num3 int) {
	fmt.Println(num3, "asdasdasdasdasd")
	if num1 == 0 {
		//vas
	} else {
		//vchs
		if num3 == 0 || num3 == 3 || num3 == 4 {
			//камера СНМ 11
			paintCHM(num1, num2, num3)
		} else {
			//камера КНК
		}
	}
	// GlobalData.Data[num1].Sub[num2].Value[num3].

	// fmt.Println(value.Name, " ", value.Desc, " ", value.val)
}

func paintCHM(num1, num2, num3 int) {
	// var ch = make(chan bool)
	// refuse(ch)
	fmt.Println(lineTop)
	fmt.Println(lineMainTop)
	for _, Fet := range Fetal {
		tempValue, _ := strconv.ParseFloat(GlobalData.Data[num1].Sub[num2].Value[num3].val, 64)
		Sf := SigF(tempValue, Fetal[0])
		fmt.Println(lineTop)
		fmt.Printf(lineStrMain, Fet, tempValue, Sf)
		// if ref := <-ch; ref == true {
		// 	continue
		// }
	}
	fmt.Println(lineTop)
}

func refuse(ch chan bool) {
	// var (
	// 	sub string
	// )
	// fmt.Scan(&sub)
	// if sub == "\n" {
	ch <- true
	// } else {
	// ch <- false
	// }
}
