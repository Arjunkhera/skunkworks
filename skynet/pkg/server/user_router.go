package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	root "skynet/pkg"

	"github.com/gorilla/mux"
)

type userRouter struct {
	userService root.UserService
}

// NewUserRouter create the router for User schema
func NewUserRouter(u root.UserService, router *mux.Router) *mux.Router {
	fmt.Println("in the router now")
	userRouter := userRouter{u}
	router.HandleFunc("/", userRouter.createUserHandler).Methods("PUT")
	router.HandleFunc("/{username}", userRouter.getUserHandler).Methods("GET")
	router.HandleFunc("/verify", userRouter.verifyUserHandler).Methods("POST")

	return router
}

func (ur *userRouter) createUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := decodeUser(r)
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = ur.userService.CreateUser(&user)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	Json(w, http.StatusOK, err)
}

func (ur *userRouter) getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	user, err := ur.userService.GetUserByUsername(username)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, user)
}

func (ur *userRouter) verifyUserHandler(w http.ResponseWriter, r *http.Request) {

	cred := root.Credentials{UserName: r.FormValue("name"), Password: r.FormValue("password")}
	res, _, flag := ur.userService.Login(cred)

	fmt.Println(res)

	if flag {
		http.Redirect(w, r, "/success", 302)
	} else {
		http.Redirect(w, r, "/failure", 302)
	}
}

// decodeUser parses the body of request
func decodeUser(r *http.Request) (root.User, error) {
	var u root.User

	if r.Body == nil {
		return u, errors.New("no request body")
	}

	err := json.NewDecoder(r.Body).Decode(&u)
	return u, err
}
