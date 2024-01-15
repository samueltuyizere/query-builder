package views

import (
	"net/http"
	"text/template"
)

func ShowMainPage(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)

}
