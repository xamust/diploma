package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/internal/app/models"
	"server/internal/app/systemsProject"
)

type Handlers struct {
	logger  *logrus.Logger
	mux     *mux.Router
	systems *systemsProject.SystemsProject
}

func (h *Handlers) handleConnection(w http.ResponseWriter, r *http.Request) {

	//for test
	resulT := &models.ResultT{}

	//incorrectdata for test...
	//if...
	resulT.Status = true
	//if...
	resulT.Error = ""
	//if...
	resulT.Data = *h.systems.GetResultData()

	nyJson, err := json.Marshal(resulT)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	if r.Method == http.MethodGet {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		w.Write(nyJson)
	}
}
