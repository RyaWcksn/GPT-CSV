package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/RyaWcksn/nann-e/api/v1/handler"
	"github.com/RyaWcksn/nann-e/api/v1/service"
	"github.com/RyaWcksn/nann-e/config"
	"github.com/RyaWcksn/nann-e/pkgs/database"
	"github.com/RyaWcksn/nann-e/pkgs/logger"
	"github.com/RyaWcksn/nann-e/store/database/prompt"
	"github.com/RyaWcksn/nann-e/store/database/user"
	"github.com/RyaWcksn/nann-e/store/gpt"
)

type Server struct {
	cfg     *config.Config
	log     logger.ILogger
	service service.IService
	handler handler.IHandler
}

var addr string
var SVR *Server
var db *sql.DB
var signalChan chan (os.Signal) = make(chan os.Signal, 1)

func (s *Server) initServer() {
	addr = ":9000"
	cfg := s.cfg
	if len(cfg.Server.HTTPAddress) > 0 {
		if _, err := strconv.Atoi(cfg.Server.HTTPAddress); err == nil {
			addr = fmt.Sprintf(":%v", cfg.Server.HTTPAddress)
		} else {
			addr = cfg.Server.HTTPAddress
		}
	}
}

func (s *Server) Register() {

	s.initServer()

	// MYSQL
	dbConn := database.NewDatabaseConnection(*s.cfg, s.log)
	if dbConn == nil {
		s.log.Fatal("Expecting DB connection but received nil")
	}

	db = dbConn.DBConnect()
	if db == nil {
		s.log.Fatal("Expecting DB connection but received nil")
	}

	user := user.NewUser(db, s.log)
	prompt := prompt.NewPrompt(db, s.log)
	openAi := gpt.NewGpt("")

	// Register service
	s.service = service.NewService(user, prompt, openAi, s.log)

	// Register handler
	s.handler = handler.NewHandler(s.service, s.log)
}

func NewService(cfg *config.Config, logger logger.ILogger) *Server {
	if SVR != nil {
		return SVR
	}
	SVR = &Server{
		cfg: cfg,
		log: logger,
	}

	SVR.Register()

	return SVR
}

func (s Server) Start() {

	go func() {
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			s.log.Fatalf("error listening to address %v, err=%v", addr, err)
		}
		s.log.Infof("HTTP server started %v", addr)
	}()

	sig := <-signalChan
	s.log.Infof("%s signal caught", sig)

	// Doing cleanup if received signal from Operating System.
	err := db.Close()
	if err != nil {
		s.log.Errorf("Error in closing DB connection. Err : %+v", err.Error())
	}
}
