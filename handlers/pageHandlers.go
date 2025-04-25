package handlers

import (
	"babybetgo/utils"
	"fmt"
	"net/http"
	"text/template"
)

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/base.htmx")
	if err != nil {
		if err != nil {
			utils.ErrorResponse(w, "base template not found", http.StatusInternalServerError)
			fmt.Println("Error parasing base template: ", err)
			return
		}
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		utils.ErrorResponse(w, "error executing base template", http.StatusInternalServerError)
		fmt.Println("Error parasing base template: ", err)
		return
	}

}

func NavbarHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/navbar.htmx")
	if err != nil {
		utils.ErrorResponse(w, "error parsing navbar.htmx tempalte", http.StatusNotFound)

	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		utils.ErrorResponse(w, "error executing the template for base.htmx: ", http.StatusInternalServerError)
	}

}
