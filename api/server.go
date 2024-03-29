package api

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/api/controllers/TEMPLATE"
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/api/controllers/health"
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/api/routes"
	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/settings"

	"github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/event/template/handler"

	"github.com/marcelofelixsalgado/financial-commons/api/middlewares"
	"github.com/marcelofelixsalgado/financial-commons/pkg/commons/logger"
	"github.com/marcelofelixsalgado/financial-commons/pkg/infrastructure/database"

	"github.com/marcelofelixsalgado/financial-commons/pkg/events"
	"github.com/marcelofelixsalgado/financial-commons/pkg/kafka"

	template "github.com/marcelofelixsalgado/financial-TEMPLATE-api/pkg/event/template"
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

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
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

	// Prometheus
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(server.http)

	// Middlewares
	server.http.Use(middlewares.Logger())

	logger.Infof("Connecting to database")

	// Connects to database
	databaseClient := database.NewConnection()

	logger.Infof("Connecting to kafka")

	// Connects to kafka
	configMap := ckafka.ConfigMap{
		"bootstrap.servers":   fmt.Sprintf("%s:%d", settings.Config.KafkaServerHost, settings.Config.KafkaServerPort),
		"delivery.timeout.ms": settings.Config.KafkaDeliveryTimeout,
		"acks":                settings.Config.KafkaAcks,
		"enable.idempotence":  settings.Config.KafkaEnableIdempotence,
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)
	eventDispatcher := events.NewEventDispatcher()

	TEMPLATERoutes := setupTEMPLATERoutes(databaseClient, eventDispatcher, kafkaProducer)
	healthRoutes := setupHealthRoutes()

	logger.Infof("Setup routes")

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

func setupTEMPLATERoutes(databaseClient *sql.DB, eventDispatcher events.IEventDispatcher, kafkaProducer *kafka.Producer) TEMPLATE.TEMPLATERoutes {
	// setup respository
	repository := TEMPLATERepository.NewTEMPLATERepository(databaseClient)

	TEMPLATECreatedEvent := template.NewTemplateCreated()
	eventDispatcher.Register("TemplateCreated", handler.NewTemplateCreatedKafkaHandler(kafkaProducer))

	// setup Use Cases (services)
	createUseCase := TEMPLATECreate.NewCreateUseCase(repository, TEMPLATECreatedEvent, eventDispatcher)
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
