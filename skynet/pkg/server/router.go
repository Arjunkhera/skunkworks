package server

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("frontPage.html", "login.html", "signup.html", "afterLogin.html", "createRecord.html"))

func renderTemplate(w http.ResponseWriter, name string) {
	err := templates.ExecuteTemplate(w, name+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createRecord(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "createRecord")
}

func frontPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "frontPage")
}

func displaySignUpHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "signup")
}

func displayLoginHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login")
}

func randomDisplay(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "afterLogin")
}
