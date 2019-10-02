package netsite

import (
	"fmt"
	"log"
	"net/http"
)

//NetVchsVas main cicle website
func NetVchsVas() {
	// http.HandleFunc("/", homePage)

	http.Handle("/", http.FileServer(http.Dir("./sitetemp")))
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "sitetemp/index.html")
	// })
	http.HandleFunc("/vchs", tableVCHS)
	http.HandleFunc("/vas", tableVAS)
	fmt.Println("Server VCHSVAS started")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("failed to start server", err)
	} else {
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprintln(writer, pageTop)
	if err != nil {
		fmt.Fprintln(writer, anError, err)
	} else {
		fmt.Fprintln(writer, mainpageinfo, buttonVas, buttonVchs)
	}
	fmt.Fprintln(writer, pageBottom)
}

func tableVAS(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprintln(writer, pageTop, `<meta http-equiv="refresh" content="0.5">`)
	fmt.Fprintln(writer, "VAS")
	if err != nil {
		fmt.Fprintln(writer, anError, err)
	} else {

		//сюда
	}
	fmt.Fprintln(writer, pageBottom)
}

func tableVCHS(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "VCHS")
}
