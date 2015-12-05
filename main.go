package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//open database
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.Handle("/", http.FileServer(http.Dir("./static/"))) //working page
	http.HandleFunc("/daily", daily)                         //upload pdf table url
	http.HandleFunc("/getCoords", getCoords)                 //ajax request to get actual markers position
	http.ListenAndServe(":4000", nil)
}

func daily(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	fmt.Println(handler.Filename)
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

func getCoords(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	GRU := Positions{[]Position{{27.926075, -15.390818},
		{27.926075, -15.390818},
		{27.926075, -15.390818},
		{27.926075, -15.390818}}}

	js, err := json.Marshal(GRU)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

type Position struct {
	Lat float32
	Lng float32
}

type Positions struct {
	Pos []Position
}

func pdf2text() {
	fileRoute := "./uploads/prueba.pdf"
	body, err := exec.Command("pdftotext", "-q", "-nopgbrk", "-enc", "UTF-8", "-eol", "unix", fileRoute, "-").Output()
	if err != nil {
		log.Println("pdf2text:", err)
	}
	//TODO: pasear contenido de body y pasarlo a un struct y este a la db
	fmt.Println(body)
}
