package server

import (
	"database/sql"
	"fmt"
	handlerusersparent "github.com/RyaWcksn/nann-e/api/v1/handler/authentication"
	handlerroles "github.com/RyaWcksn/nann-e/api/v1/handler/roles"
	handlerchild "github.com/RyaWcksn/nann-e/api/v1/handler/user_child"
	serviceusersparent "github.com/RyaWcksn/nann-e/api/v1/service/authentication"
	serviceroles "github.com/RyaWcksn/nann-e/api/v1/service/roles"
	servicechild "github.com/RyaWcksn/nann-e/api/v1/service/user_child"
	"github.com/RyaWcksn/nann-e/pkgs/database/mysql"
	"github.com/RyaWcksn/nann-e/server/middleware"
	storeroles "github.com/RyaWcksn/nann-e/store/database/roles"
	storeusersparent "github.com/RyaWcksn/nann-e/store/database/user"
	storechild "github.com/RyaWcksn/nann-e/store/database/user_child"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"strconv"

	"github.com/RyaWcksn/nann-e/config"
	"github.com/RyaWcksn/nann-e/pkgs/logger"
)

type Server struct {
	cfg *config.Config
	log logger.ILogger

	// Users Parent
	serviceUsersParent serviceusersparent.IService
	handlerUsersParent handlerusersparent.IHandler

	// Roles
	serviceRoles serviceroles.IService
	handlerRoles handlerroles.IHandler

	// Child
	serviceChild servicechild.IService
	handlerChild handlerchild.IHandler
}

var addr string
var SVR *Server
var db *sql.DB
var signalChan chan (os.Signal) = make(chan os.Signal, 1)
var ViberApp *fiber.App

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
	dbConn := mysql.NewDatabaseConnection(*s.cfg, s.log)
	if dbConn == nil {
		s.log.Fatal("Expecting DB connection but received nil")
	}

	db = dbConn.DBConnect()
	if db == nil {
		s.log.Fatal("Expecting DB connection but received nil")
	}

	usersParentRepo := storeusersparent.NewUserParentImpl(db, s.log)
	rolesRepo := storeroles.NewRolesImpl(db, s.log)
	childRepo := storechild.NewChildImpl(db, s.log)

	// Register service
	s.serviceUsersParent = serviceusersparent.NewServiceImpl(usersParentRepo, s.cfg, s.log)
	s.serviceRoles = serviceroles.NewRolesService(rolesRepo, s.log)
	s.serviceChild = servicechild.NewChildService(childRepo, s.log)

	// Register handler
	s.handlerUsersParent = handlerusersparent.NewUsersParentHandler(s.serviceUsersParent, s.log)
	s.handlerRoles = handlerroles.NewRoles(s.serviceRoles, s.log)
	s.handlerChild = handlerchild.NewChildHandler(s.serviceChild, s.log)
}

func New(cfg *config.Config, logger logger.ILogger) *Server {
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
	ViberApp = fiber.New(fiber.Config{
		Immutable: true,
	})

	// API AUTHENTICATION
	auth := ViberApp.Group("/api/v1/auth")
	auth.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	auth.Use(middleware.ErrorHandler)

	// authentication
	auth.Post("/user/register", s.handlerUsersParent.RegisterParent)
	auth.Post("/user/login", s.handlerUsersParent.LoginParent)

	// BUSINESS API
	v1 := ViberApp.Group("/api/v1")
	v1.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	v1.Use(middleware.Authenticate(s.cfg, s.log))
	v1.Use(middleware.ErrorHandler)

	// roles
	v1.Post("/roles", s.handlerRoles.CreateRoles)
	v1.Get("/role/:roleName", s.handlerRoles.GetOneRoleById)
	v1.Get("/roles", s.handlerRoles.GetListRole)
	v1.Patch("/role/:roleName", s.handlerRoles.UpdateSingleRole)

	// Child
	v1.Post("/child", s.handlerChild.CreateUserChild)

	go func() {
		err := ViberApp.Listen(":9000")
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
