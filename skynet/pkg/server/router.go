package server

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("login.html", "success.html", "failure.html"))

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

func displayLoginHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login")
}

func verifyLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("we were here")
	http.Redirect(w, r, "/user/verify", 302)
}
