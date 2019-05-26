package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	root "skynet/pkg"

	"github.com/gorilla/mux"
)

type claimRouter struct {
	claimService root.ClaimService
	port         string
}

// NewRecordRouter create the router for Record schema
func NewClaimRouter(claim root.ClaimService, router *mux.Router, port string) *mux.Router {
	claimrt := claimRouter{claim, port}

	router.HandleFunc("/create", claimrt.createClaimHandler)
	router.HandleFunc("/displayAllClaims", claimrt.displayAllClaims)
	router.HandleFunc("/displayAllClaimDefns", claimrt.displayAllClaimDefns)
	//router.HandleFunc("/displayAll", recordRouter.displayAllRecords).Methods("GET")

	return router
}

func (claim *claimRouter) createClaimHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		createClaim(w, r)
		return
	}

	fmt.Println("1")
	result := make(map[string]string)

	for i := 1; i < 4; i++ {
		fmt.Println("entry")
		result[r.FormValue("attr"+strconv.Itoa(i))] = r.FormValue("type" + strconv.Itoa(i))
		fmt.Println(r.FormValue("attr" + strconv.Itoa(1)))
	}

	fmt.Println("2")
	identifier, err := claim.claimService.CreateClaimDefn(result)
	usr, err := http.Get("http://localhost" + claim.port + "/user/arjun")

	var u root.User
	err = json.NewDecoder(usr.Body).Decode(&u)

	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("3")
	claim.claimService.CreateClaim(u.Identifier, identifier)
	//http.Redirect(w, r, "/", 302)
	http.Redirect(w, r, "/claim/displayAllClaims", 302)
}

func (claimrt *claimRouter) displayAllClaims(w http.ResponseWriter, r *http.Request) {

	results, err := claimrt.claimService.GetAllClaims()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	bytes, err := json.MarshalIndent(results, "", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	io.WriteString(w, string(bytes))
}

func (claimrt *claimRouter) displayAllClaimDefns(w http.ResponseWriter, r *http.Request) {

	results, err := claimrt.claimService.GetAllClaimDefns()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	bytes, err := json.MarshalIndent(results, "", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	io.WriteString(w, string(bytes))
}

/*
func (rec *recordRouter) displayAllRecords(w http.ResponseWriter, r *http.Request) {

	results, err := rec.recordService.GetAllRecords()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	bytes, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		http.Error(w, err.Erro(), http.StatusInternalServerError)
	}

	io.WriteString(w, string(bytes))
}

*/
