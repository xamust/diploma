package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"server/internal/app/collect"
	"server/internal/app/systemsProject"
	"server/testing/emulator"
	"sync"
)

type AppServer struct {
	config  *Config
	mux     *mux.Router
	logger  *logrus.Logger
	handl   Handlers
	collect *collect.Collect
	systems *systemsProject.SystemsProject
	wg      sync.WaitGroup
	mu      sync.Mutex
}

// init new server
func New(config *Config) *AppServer {
	return &AppServer{
		config: config,
		mux:    mux.NewRouter(),
		logger: logrus.New(),
	}
}

// configure logrus...
func (s *AppServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	s.logger.Info("Логгер инициализирован успешно!")
	return nil
}

// configure emulator
func (s *AppServer) configureEmulator() {
	go emulator.EmulatorMain()
	s.logger.Info("Эмулятор запущен успешно!")
}

// config route...
func (s *AppServer) configureRouter() {
	s.mux.HandleFunc("/", s.handl.handleConnection)
	//serve CSS files with gorilla mux...
	s.mux.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))
	s.mux.Handle("/status_page.html", http.FileServer(http.Dir("./web")))
	s.logger.Info("Gorilla mux инициализирован успешно!")
}

// config collecting service...
func (s *AppServer) configureCollect() error {
	s.collect = &collect.Collect{Logger: s.logger, Config: s.config.Collect}
	s.logger.Info("Коллектор *.data файлов инициализирован успешно!")
	return s.collect.Start()
}

// config systems....
func (s *AppServer) configureSystems() {
	s.systems = &systemsProject.SystemsProject{ParsingDataFiles: s.collect.ParsingDataFiles, Config: s.config.Systems}
	s.logger.Info("Системы инициализированы успешно!")
}

func (s *AppServer) Start() error {

	//configure emulator...
	s.configureEmulator()

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
	go s.configureRouter()

	//handlers init...
	s.handl = Handlers{s.logger, s.mux, s.systems}
	s.logger.Info(fmt.Sprintf("Starting server (bind on %v)...", s.config.BindAddr)) // set message Info level about succesfull starting server...
	//for heroku
	if os.Getenv("PORT") != "" {
		s.config.BindAddr = fmt.Sprintf(":%s", os.Getenv("PORT"))
	}

	return http.ListenAndServe(s.config.BindAddr, s.mux) //bind addr from Config and new gorilla mux
}
