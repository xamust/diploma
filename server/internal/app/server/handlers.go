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
	data, err := h.systems.GetResultData()

	if err != nil {
		resulT.Error = err.Error()
		resulT.Status = false
		resulT.Data = models.ResultSetT{}
	} else {
		resulT.Error = ""
		resulT.Status = true
		resulT.Data = *data
	}

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
