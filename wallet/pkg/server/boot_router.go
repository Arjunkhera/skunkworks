package server

import (
	"net/http"
	root "wallet/pkg"

	"github.com/gorilla/mux"
)

type bootRouter struct {
	bootService   root.UserService
	deviceService root.DeviceService
}

// NewBootRouter create the router for User schema
func NewBootRouter(u root.UserService, d root.DeviceService, router *mux.Router) *mux.Router {

	btRouter := bootRouter{u, d}
	router.HandleFunc("/create", btRouter.createUserHandler).Methods("POST")
	router.HandleFunc("/verify", btRouter.verifyUserHandler).Methods("POST")

	return router
}

func (ur *bootRouter) createUserHandler(w http.ResponseWriter, r *http.Request) {

	cred := root.User{UserName: r.FormValue("name"), Password: r.FormValue("password")}
	// write the credentials to the boot file
	uniqueID, err := ur.bootService.CreateUser(&cred)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	dev := root.Device{Identifier: uniqueID, PublicKey: ""}
	err = ur.deviceService.CreateDevice(&dev)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/view", 302)
}

func (ur *bootRouter) verifyUserHandler(w http.ResponseWriter, r *http.Request) {

	cred := root.User{UserName: r.FormValue("name"), Password: r.FormValue("password")}
	// verify the credentials from the boot file
	flag, err := ur.bootService.Login(cred)
	if !flag {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/view", 302)
	return
}
