package api

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/api/controllers/TEMPLATE"
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/api/controllers/health"
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/api/routes"
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/settings"
	"github.com/marcelofelixsalgado/financial-commons/api/middlewares"
	"github.com/marcelofelixsalgado/financial-commons/pkg/commons/logger"
	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/database"

	TEMPLATERepository "github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/infrastructure/repository/TEMPLATE"
	TEMPLATECreate "github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/usecase/TEMPLATE/create"
	TEMPLATEDelete "github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/usecase/TEMPLATE/delete"
	TEMPLATEFind "github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/usecase/TEMPLATE/find"
	TEMPLATEList "github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/usecase/TEMPLATE/list"
	TEMPLATEUpdate "github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/usecase/TEMPLATE/update"

	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

// Server this is responsible for running an http server
type Server struct {
	http   *echo.Echo
	routes *routes.Routes
	stop   chan struct{}
}

// Start - Entry point of the API
func NewServer() *Server {
	// Load environment variables
	settings.Load()

	server := &Server{
		stop: make(chan struct{}),
	}

	return server
}

// Run is the procedure main for start the application
func (s *Server) Run() {
	s.startServer()
	<-s.stop
}

func (server *Server) startServer() {
	go server.watchStop()

	server.http = echo.New()
	logger := logger.GetLogger()

	logger.Infof("Server is starting now in %s.", settings.Config.Environment)

	// Middlewares
	server.http.Use(middlewares.Logger())

	// Connects to database
	databaseClient := database.NewConnection()

	TEMPLATERoutes := setupTEMPLATERoutes(databaseClient)
	healthRoutes := setupHealthRoutes()

	// Setup all routes
	routes := routes.NewRoutes(TEMPLATERoutes, healthRoutes)

	routes.RouteMapping(server.http)
	server.routes = routes

	showRoutes(server.http)

	addr := fmt.Sprintf(":%v", settings.Config.ApiHttpPort)
	go func() {
		if err := server.http.Start(addr); err != nil {
			logger.Errorf("Shutting down the server now: %s", err)
		}
	}()
}

// watchStop wait for the interrupt signal.
func (server *Server) watchStop() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	logger.GetLogger().Info(<-stop)
	server.stopServer()
}

// stopServer stops the server http
func (s *Server) stopServer() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(settings.Config.ServerCloseWait))
	defer cancel()

	logger := logger.GetLogger()
	logger.Info("Server is stoping...")
	s.http.Shutdown(ctx)
	close(s.stop)
}

func setupTEMPLATERoutes(databaseClient *sql.DB) TEMPLATE.TEMPLATERoutes {
	// setup respository
	repository := TEMPLATERepository.NewTEMPLATERepository(databaseClient)

	// setup Use Cases (services)
	createUseCase := TEMPLATECreate.NewCreateUseCase(repository)
	deleteUseCase := TEMPLATEDelete.NewDeleteUseCase(repository)
	findUseCase := TEMPLATEFind.NewFindUseCase(repository)
	listUseCase := TEMPLATEList.NewListUseCase(repository)
	updateUseCase := TEMPLATEUpdate.NewUpdateUseCase(repository)

	// setup router handlers
	TEMPLATEHandler := TEMPLATE.NewTEMPLATEHandler(createUseCase, deleteUseCase, findUseCase, listUseCase, updateUseCase)

	// setup routes
	TEMPLATERoutes := TEMPLATE.NewTEMPLATERoutes(TEMPLATEHandler)

	return TEMPLATERoutes
}

func setupHealthRoutes() health.HealthRoutes {
	// setup router handlers
	healthHandler := health.NewHealthHandler()

	// setup routes
	healthRoutes := health.NewHealthRoutes(healthHandler)

	return healthRoutes
}

func showRoutes(e *echo.Echo) {
	var routes = e.Routes()
	logger := logger.GetLogger()

	if len(routes) > 0 {
		for _, route := range routes {
			logger.Infof("%6s: %s \n", route.Method, route.Path)
		}
	}
}
