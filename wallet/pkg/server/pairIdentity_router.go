package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	root "wallet/pkg"

	"github.com/gorilla/mux"
)

type pairIdentityRouter struct {
	pairIdentityService root.PairIdentityService
	SkynetIP            string
}

// NewPairIdentityRouter create the router for PairIdentity schema
func NewPairIdentityRouter(pId root.PairIdentityService, router *mux.Router, skynetIP string) *mux.Router {

	pIdRouter := pairIdentityRouter{pId, skynetIP}
	router.HandleFunc("/create", pIdRouter.createPairIdentityHandler).Methods("POST")
	router.HandleFunc("/all", pIdRouter.getAllPairIdentitiesHandler)
	router.HandleFunc("/obtainClaim", pIdRouter.getClaimHandler)
	// router.HandleFunc("/{pId}", pIdRouter.getPairIdentityHandler).Methods("GET")
	// router.HandleFunc("/verify", btRouter.verifyUserHandler).Methods("POST")

	return router
}

func (pId *pairIdentityRouter) createPairIdentityHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("IdentityName")
	otherPartyName := r.FormValue("OtherParty")

	err := pId.pairIdentityService.CreatePairIdentity(username, otherPartyName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	userID, _ := pId.pairIdentityService.GetPairIdentityByUsername(username)

	fmt.Println("the middle")

	var url string = "http://192.168.1.3:8091/IdentityChain"
	type chain struct {
		DID       string `json:"DID"`
		PublicKey string `json:"Public Key"`
		WalletID  string `json:"Wallet ID"`
	}

	temp := chain{
		DID:       userID.Identifier,
		PublicKey: userID.PublicKey,
		WalletID:  "randomString"}

	msg, _ := json.Marshal(temp)
	/*
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(msg))

		fmt.Println(temp)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, _ := client.Do(req)

		defer resp.Body.Close()
	*/
	http.Post(url, "application/json", bytes.NewBuffer(msg))

	fmt.Println("the end")
	//http.Redirect(w, r, "/view", 302)
}

func (pId *pairIdentityRouter) getPairIdentityHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pIdname := vars["pId"]

	_, err := pId.pairIdentityService.GetPairIdentityByUsername(pIdname)
	//_, err := pId.pairIdentityService.GetPairIdentityByUsername(pIdname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/pairId/all", 200)
}

func (pId *pairIdentityRouter) getClaimHandler(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, pId.SkynetIP+"/claim/getClaimDefn", 302)

}

func (pId *pairIdentityRouter) getAllPairIdentitiesHandler(w http.ResponseWriter, r *http.Request) {

	results, err := pId.pairIdentityService.GetAllPairIdentities()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	bytes, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	io.WriteString(w, string(bytes))
}
