package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"
)

var DB *sql.DB // Make sure this is initialized in your main/server setup

// GetUserBalanceHandler returns the user's balance as a snippet for htmx polling
func GetUserBalanceHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Replace with actual user identification (e.g., from session/cookie)
	userID := 1
	var balance float64
	err := DB.QueryRow("SELECT balance FROM users WHERE id=$1", userID).Scan(&balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `<span id=\"balance-value\">Error</span>`)
		return
	}
	fmt.Fprintf(w, `<span id=\"balance-value\" hx-get=\"/get_user_balance\" hx-trigger=\"every 1s\" hx-swap=\"outerHTML\">$%.2f</span>`, balance)
}

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
