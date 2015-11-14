package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static/"))) //working page
	http.HandleFunc("/daily", daily)                         //upload pdf table url
	//http.HandleFunc("/getCoords", getCoords)	//ajax request to get actual markers position
	http.ListenAndServe(":4000", nil)
}

func daily(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	fmt.Print(handler.Filename)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	//err = ioutil.WriteFile(handler.Filename, data, 0777)
	err = ioutil.WriteFile("uploads/"+handler.Filename, data, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

// func getCoords(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.NotFound(w, r)
// 		return
// 	}
//
// 	field := r.FormValue("textfield")
// 	mark := markdown.NewParser(strings.NewReader(field))
// 	log.Println("Ajax:", field)
// 	w.Write([]byte(mark.Markdown()))
// }
