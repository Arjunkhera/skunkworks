package server

import (
	"net/http"

	root "skynet/pkg"

	"github.com/gorilla/mux"
)

type recordRouter struct {
	recordService root.RecordService
}

// NewRecordRouter create the router for Record schema
func NewRecordRouter(rec root.RecordService, router *mux.Router) *mux.Router {
	recordRouter := recordRouter{rec}

	router.HandleFunc("/create", recordRouter.createRecordHandler).Methods("POST")

	return router
}

func (rec *recordRouter) createRecordHandler(w http.ResponseWriter, r *http.Request) {
	Record := root.Record{CommonName: r.FormValue("commonName")}
	http.Redirect(w, r, r.FormValue, 302)

	err := rec.recordService.CreateRecord(&Record, r.FormValue("name"))
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.Redirect(w, r, "/display", 302)
}
