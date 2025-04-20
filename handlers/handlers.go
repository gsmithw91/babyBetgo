package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"
)

var DB *sql.DB // Make sure this is initialized in your main/server setup

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(
		"templates/index.htmx",
		"templates/auth_modal.htmx",
		"templates/balance.htmx",
		"templates/events.htmx",
		"templates/bet_form.htmx",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error reading templates:", err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error executing template:", err)
	}

}
