package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/misuher/RoutesMap/Coordinates"
	"github.com/misuher/RoutesMap/models"
)

type Env struct {
	db models.Datastore
}

func main() {
	db, err := models.Open()
	if err != nil {
		log.Panic(err)
	}
	models.Create()
	models.CreateParking()
	env := &Env{db}

	http.Handle("/", http.FileServer(http.Dir("./static/"))) //working page
	http.HandleFunc("/daily", env.dailyUpload)               //upload pdf table url
	http.HandleFunc("/getCoords", env.getCoords)             //ajax request to get actual markers position
	http.ListenAndServe(":4000", nil)
}

func (env *Env) dailyUpload(w http.ResponseWriter, r *http.Request) {
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
	//dynamic file name
	var fileRoute bytes.Buffer
	fileRoute.WriteString("./uploads/")
	fileRoute.WriteString(time.Now().Format("01/02/2006 03:04:05"))
	fileRoute.WriteString(".pdf")
	err = ioutil.WriteFile(fileRoute.String(), data, 0777)
	if err != nil {
		fmt.Println(err)
	}
	last.setLastFile(fileRoute.String())
	pdf2text()
}

func (env *Env) getCoords(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	//TODO: calculate coords dynamically
	markers := coord.Positions{
		[]coord.Position{
			{27.926075, -15.390818},
			{27.926075, -15.390818},
			{27.926075, -15.390818},
			{27.926075, -15.390818},
		}}

	js, err := json.Marshal(markers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func pdf2text() error {
	//convert pdf to text
	body, err := exec.Command("pdftotext", "-q", "-nopgbrk", "-enc", "UTF-8", "-eol", "unix", last.getLastFile(), "-").Output()
	if err != nil {
		return err
	}

	//TODO: parsear contenido de body y pasarlo a un struct y este a la db
	fmt.Println(body)
	return nil
}

type lastFile struct {
	fileName string
}

func (l *lastFile) setLastFile(filename string) {
	l.fileName = filename
}

func (l *lastFile) getLastFile() string {
	return l.fileName
}

var last lastFile
