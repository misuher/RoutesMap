package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static/"))) //working page
	http.HandleFunc("/coordinates", coordinates)                     //ajax url
	http.ListenAndServe(":4000", nil)
}

func coordinates(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	field := r.FormValue("textfield")
	mark := markdown.NewParser(strings.NewReader(field))
	log.Println("Ajax:", field)
	w.Write([]byte(mark.Markdown()))
}
