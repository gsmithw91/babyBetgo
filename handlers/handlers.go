package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	templateDir := "templates/index.htmx"
	tmpl, err := template.ParseFiles(templateDir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error reading template at :", templateDir)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error reading template at :", templateDir)
	}

}
