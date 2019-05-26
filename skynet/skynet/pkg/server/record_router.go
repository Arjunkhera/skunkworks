package server

/*
import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"skynet/pkg"
)

type recordRouter struct {
	recordService root.RecordService
}

func NewRecordRouter(rec root.RecordService, router *mux.Router) *mux.Router {
	recordRouter := recordRouter{rec}
	router.HandleFunc("/").Methods("GET")
	router.HandleFunc("/new", recordRouter.createRouterHandler).Methods("PUT")

	return router
}

func (rec *recordRouter) createRouterHandler(w http.ResponseWriter, r *http.Request) {
	err, user := decodeRouter(r)
}

func (rec *recordRouter) verififyUser() {
}

func decodeRouter(r *http.Request) (root.Router, error) {

}
*/
