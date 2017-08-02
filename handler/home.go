package handler

import (
	"net/http"
	"github.com/onezerobinary/MyFirstGoApp/model"
	"html/template"
)

var home = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))


func HomeHandler(w http.ResponseWriter, req *http.Request) {

	enrico := model.Person{
		Name:"Enrico",
		Surname: "Zanardo",
		Email:"ezanardo@onezerobinary.com",
	}

	home.Execute(w, enrico)
}
