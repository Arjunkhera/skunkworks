package server

import (
	"net/http"

	root "skynet/pkg"

	"github.com/gorilla/mux"
)

type pairIdentityRouter struct {
	pairIdentityService root.PairIdentityService
	port                string
}

// NewRecordRouter create the router for Record schema
func NewPairIdentityRouter(pairId root.PairIdentityService, router *mux.Router, port string) *mux.Router {
	pairrt := pairIdentityRouter{pairId, port}

	router.HandleFunc("/create", pairrt.createPairIdentityHandler).Methods("POST")
	// router.HandleFunc("/displayAllClaims", claimrt.displayAllClaims)
	// router.HandleFunc("/displayAllClaimDefns", claimrt.displayAllClaimDefns)

	return router
}

func (pairId *pairIdentityRouter) createPairIdentityHandler(w http.ResponseWriter, r *http.Request) {

	//err := json.NewDecoder(r.Body).Decode(&u)
}

/*
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
*/
