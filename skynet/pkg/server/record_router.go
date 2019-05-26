package server

import (
	"encoding/json"
	"net/http"

	root "skynet/pkg"

	"github.com/gorilla/mux"
)

type recordRouter struct {
	recordService root.RecordService
	port          string
}

// NewRecordRouter create the router for Record schema
func NewRecordRouter(rec root.RecordService, router *mux.Router, port string) *mux.Router {
	recordRouter := recordRouter{rec, port}

	router.HandleFunc("/create", recordRouter.createRecordHandler).Methods("POST")

	return router
}

func (rec *recordRouter) createRecordHandler(w http.ResponseWriter, r *http.Request) {
	Record := root.Record{CommonName: r.FormValue("commonName")}
	usr, err := http.Get("localhost" + rec.port + "/user/" + r.FormValue("name"))

	var u root.User
	err := json.NewDecoder(usr.Body).Decode(&u)
	Record.Identifier = u.Identifier
	err := rec.recordService.CreateRecord(&Record)

	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.Redirect(w, r, "/display", 302)
}
