package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/internal/app/collect"
	"server/internal/app/systemsProject"
)

type AppServer struct {
	config  *Config
	mux     *mux.Router
	logger  *logrus.Logger
	handl   Handlers
	collect *collect.Collect
	systems *systemsProject.SystemsProject
}

//init new server
func New(config *Config) *AppServer {
	return &AppServer{
		config: config,
		mux:    mux.NewRouter(), //gorilla/mux
		logger: logrus.New(),    //sirupsen/logrus
	}
}

//configure logrus...
func (s *AppServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *AppServer) Start() error {

	//configure logger...
	if err := s.configureLogger(); err != nil {
		return err //if logrus configure result err
	}

	//configure collecting...
	if err := s.configureCollect(); err != nil {
		return err
	}
	//configure systemsProject...
	s.configureSystems()

	//configure router...
	s.configureRouter()

	//handlers init...
	s.handl = Handlers{s.logger, s.mux, s.systems}
	s.logger.Info(fmt.Sprintf("Starting server (bind on %v)...", s.config.BindAddr)) // set message Info level about succesfull starting server...
	return http.ListenAndServe(s.config.BindAddr, s.mux)                             //bind addr from Config and new gorilla mux
}

//config route...
func (s *AppServer) configureRouter() {
	s.mux.HandleFunc("/", s.handl.handleConnection)
}

//config collecting service...
func (s *AppServer) configureCollect() error {
	s.collect = &collect.Collect{Logger: s.logger, Config: s.config.Collect}
	return s.collect.Start()
}

//config systems....
func (s *AppServer) configureSystems() {
	s.systems = &systemsProject.SystemsProject{ParsingDataFiles: s.collect.ParsingDataFiles, Config: s.config.Systems}
}
