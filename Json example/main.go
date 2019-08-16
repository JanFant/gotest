// Работа с Json
package main

import (
	"encoding/json"
	"fmt"
)

//Test эксперементы с json
type Test struct {
	Aaa string `json:"a"`
	Bbb int    `json:"b"`
	Ll  int64  `json:""`
	Zz  bool
}

func main() {
	fmt.Println("Ex1")
	a := Test{Aaa: "asd", Bbb: 1, Ll: 4, Zz: true}
	data, _ := json.Marshal(&a)
	fmt.Println(string(data))

	sliceVar1 := []string{"John", "Andrew", "Robert"}
	sliceVar2, _ := json.Marshal(sliceVar1)
	fmt.Println(string(sliceVar2))
	fmt.Println(sliceVar1)

	mapVar1 := map[string]string{"John": "Accepted", "Andrew": "Waiting", "Robert": "Cancelled"}
	mapVar2, _ := json.Marshal(mapVar1)
	fmt.Println(string(mapVar2))
	fmt.Println(mapVar1)

	fmt.Println("Ex2")
	fmt.Println("----------------------------")

	// якобы пишедший json
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	//для первого элемента явно указываем тип инт
	num := dat["num"].(float64) // для того, чтобы получить из свойства num число
	fmt.Println(num)

	strs := dat["strs"].([]interface{}) // для того, чтобы получить массив интерфейсов...
	fmt.Println(strs)
	// для второго массива и каждого элемента явно указываем что это строка
	str1 := strs[0].(string) // ... и потом получить из него строку
	fmt.Println(str1)

	fmt.Println("Ex3")
	fmt.Println("----------------------------")

	user := User{}
	userJson := "{\"FirstName\":\"John\",\"LastName\":\"Smith\",\"Books\":[\"The Art of Programming\",\"Golang for Dummies\"]}"
	bytes := []byte(userJson)
	json.Unmarshal(bytes, &user)
	fmt.Println(user.FirstName, user.LastName, user.Books)
	// John Smith [The Art of Programming Golang for Dummies]

	user2 := User2{}
	userJson2 := "{\"name\":\"John\",\"lastname\":\"Smith\",\"ordered_books\":[\"The Art of Programming\",\"Golang for Dummies\"]}"
	bytes2 := []byte(userJson2)
	json.Unmarshal(bytes2, &user2)
	fmt.Println(user2.FirstName, user2.LastName, user2.Books)
	// John Smith [The Art of Programming Golang for Dummies]
}

// User bla bla
type User struct {
	FirstName string
	LastName  string
	Books     []string
}

//User2 bla bla bla
type User2 struct {
	FirstName string   `json:"name"`          // свойство FirstName будет преобразовано в ключ "name"
	LastName  string   `json:"lastname"`      // свойство LastName будет преобразовано в ключ "lastname"
	Books     []string `json:"ordered_books"` // свойство Books будет преобразовано в ключ "ordered_books"
}
