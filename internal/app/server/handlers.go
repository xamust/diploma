package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/internal/app/casher"
	"server/internal/app/systemsproject"
)

type Handlers struct {
	logger  *logrus.Logger
	mux     *mux.Router
	systems *systemsproject.SystemsProject
	casher  *casher.Casher
}

func (h *Handlers) handleConnection(w http.ResponseWriter, r *http.Request) {

	nyJSON, err := json.Marshal(h.casher.ToHandler())
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		w.Write(nyJSON)
	}
}
