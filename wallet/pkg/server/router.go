package server

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("login.html", "signup.html", "success.html", "failure.html", "afterLogin.html"))

func renderTemplate(w http.ResponseWriter, name string) {
	err := templates.ExecuteTemplate(w, name+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func displaySignUpHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "signup")
}

func displayLoginHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login")
}

func displayViewHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "afterLogin")
}
