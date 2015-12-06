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
	"github.com/misuher/RoutesMap/parser"
)

type Env struct {
	db models.Datastore
}

func main() {
	db, err := models.Open()
	if err != nil {
		log.Panic(err)
	}

	db.Create()
	db.CreateParking()
	env := &Env{db: db}

	http.Handle("/", http.FileServer(http.Dir("./static/"))) //working page
	http.HandleFunc("/daily", env.dailyUpload)               //upload pdf table url
	http.HandleFunc("/getCoords", env.getCoords)             //ajax request to get actual markers position
	http.ListenAndServe(":4000", nil)
}

func (env *Env) dailyUpload(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	fmt.Println(handler.Filename)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}
	//err = ioutil.WriteFile(handler.Filename, data, 0777)
	//dynamic file name
	var fileRoute bytes.Buffer
	fileRoute.WriteString("./uploads/")
	fileRoute.WriteString(time.Now().Format("2006-01-02"))
	fileRoute.WriteString(".pdf")
	err = ioutil.WriteFile(fileRoute.String(), data, 0777)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	err = env.pdf2text()
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	//for _, bk := range bks {
	//	fmt.Fprintf(w, "%s %s  \n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	//}
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

func (env *Env) pdf2text() error {
	var fileRoute bytes.Buffer
	fileRoute.WriteString("./uploads/")
	fileRoute.WriteString(time.Now().Format("2006-01-02"))
	fileRoute.WriteString(".pdf")
	//convert pdf to text
	body, err := exec.Command("pdftotext", "-q", "-nopgbrk", "-enc", "UTF-8", "-eol", "unix", fileRoute.String(), "-").Output()
	if err != nil {
		log.Println("No hay daily de hoy")
		return err
	}
	date, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	err = env.db.DeleteDaily(date)
	err = parser.ParsePDF(body, env.db)
	if err != nil {
		return err
	}
	return nil
}
