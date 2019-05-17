package server

import (
	"net/http"
	root "wallet/pkg"

	"github.com/gorilla/mux"
)

type bootRouter struct {
	bootService root.UserService
}

// NewBootRouter create the router for User schema
func NewBootRouter(u root.UserService, router *mux.Router) *mux.Router {

	btRouter := bootRouter{u}
	router.HandleFunc("/create", btRouter.createUserHandler).Methods("POST")
	router.HandleFunc("/verify", btRouter.verifyUserHandler).Methods("POST")

	return router
}

func (ur *bootRouter) createUserHandler(w http.ResponseWriter, r *http.Request) {

	cred := root.User{UserName: r.FormValue("name"), Password: r.FormValue("password")}
	// write the credentials to the boot file
	err := ur.bootService.CreateUser(&cred)
	if err != nil {
		http.Redirect(w, r, "/failure", 302)
	}

	http.Redirect(w, r, "/view", 302)
}

func (ur *bootRouter) verifyUserHandler(w http.ResponseWriter, r *http.Request) {

	cred := root.User{UserName: r.FormValue("name"), Password: r.FormValue("password")}
	// verify the credentials from the boot file
	flag, _ := ur.bootService.Login(cred)
	if !flag {
		http.Redirect(w, r, "/failure", 302)
	}

	http.Redirect(w, r, "/view", 302)
}
