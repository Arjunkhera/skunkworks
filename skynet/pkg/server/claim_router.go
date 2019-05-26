package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	root "skynet/pkg"

	"fmt"
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
	//router.HandleFunc("/displayAll", recordRouter.displayAllRecords).Methods("GET")

	return router
}

func (claim *claimRouter) createClaimHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		createClaim(w, r)
		return
	}

	var result map[string]string

	for i := 0; i < 3; i++ {
		result[r.FormValue("attr"+strconv.Itoa(i))] = r.FormValue("" + strconv.Itoa(i))
	}

	identifier, err := claim.claimService.CreateClaimDefn(result)
	usr, err := http.Get("http://localhost" + claim.port + "/user/" + r.FormValue("name"))

	var u root.User
	err = json.NewDecoder(usr.Body).Decode(&u)

	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	claim.claimService.CreateClaim(u.Identifier, identifier)

}

func (claimrt *claimRouter) displayAllClaims(w http.ResponseWriter, r *http.Request) {

	results, err := claimrt.claimService.GetAllClaims()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	bytes, err := json.MarshalIdent(results, "", "")
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

	bytes, err := json.MarshalIdent(results, "", "")
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
