package server

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("login.html", "signup.html", "success.html", "failure.html", "randomview.html"))

func renderTemplate(w http.ResponseWriter, name string) {
	err := templates.ExecuteTemplate(w, name+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func itWorked(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "success")
}

func itDidNotWork(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "failure")
}

func displaySignUpHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "signup")
}

func displayLoginHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login")
}

func randomViewHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "randomview")
}
