package server

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("frontPage.html", "login.html", "signup.html", "randomDisplay.html", "createClaim.html"))

func renderTemplate(w http.ResponseWriter, name string) {
	err := templates.ExecuteTemplate(w, name+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createClaim(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "createClaim")
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
	renderTemplate(w, "randomDisplay")
}

/*
func itWorked(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "success")
}

func itDidNotWork(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "failure")
}

func verifyLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("we were here")
	http.Redirect(w, r, "/user/verify", 302)
}
*/
