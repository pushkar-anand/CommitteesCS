package main

import (
	"committees/config"
	"committees/db"
	"committees/handlers"
	"committees/security"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Server wraps server data and functions
type Server struct {
	server http.Server

	router      *mux.Router
	HTTPHandler http.Handler
	logger      *logrus.Logger
	appConfig   *config.AppConfig
	killServer  chan int
	connClose   chan int
	db          *db.DB
}

// NewServer creates an instance of the Server
func NewServer(logger *logrus.Logger, appConfig *config.AppConfig) *Server {
	server := &Server{
		logger:     logger,
		appConfig:  appConfig,
		killServer: make(chan int),
		connClose:  make(chan int),
	}

	r := mux.NewRouter()
	r.StrictSlash(true)

	server.router = r

	return server
}

// Initialize the Server
func (s *Server) Initialize() {
	s.connectDB()
	s.addMiddleWares()
	s.addRoutes()
	s.addHandlers()

	// This runs in background and listens for any signal from the OS
	go func() {
		sigint := make(chan os.Signal)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGKILL)

		select {
		case sig := <-sigint:
			s.logger.Infof("Shutdown signal received: %v", sig)
		case <-s.killServer:
			s.logger.Info("Kill server request")
		}

		s.logger.Debug("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		err := s.server.Shutdown(ctx)
		if err != nil {
			s.logger.WithError(err).Error("Error shutting down server")
		}
		close(s.connClose)
	}()

	// To ensure graceful shutdown when a fatal log is logged
	s.logger.ExitFunc = func(code int) {
		s.logger.Info("Issuing kill request")
		s.killServer <- 1
	}
}

func (s *Server) connectDB() {
	s.db = db.GetDB(s.logger)
}

func (s *Server) addRoutes() {
	addRoutes(s.router, s.logger)
}

func (s *Server) addMiddleWares() {
	//s.router.Use(validation.UUIDValidator(s.logger, "id", "uuid"))
}

func (s *Server) addHandlers() {
	s.HTTPHandler = s.router

	s.HTTPHandler = handlers.NewLoggingHandler(s.logger)(s.HTTPHandler)
	s.HTTPHandler = security.SecureHandler(s.appConfig.Production)(s.HTTPHandler)
	s.HTTPHandler = handlers.RecoveryHandler(s.logger, s.appConfig.Production)(s.HTTPHandler)
	s.HTTPHandler = handlers.AssignRequestIDHandler(s.HTTPHandler)
}

// Listen starts the server
func (s *Server) Listen() {
	addr := fmt.Sprintf("%s:%d", "", s.appConfig.PORT)

	s.server = http.Server{
		Addr:         addr,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      s.HTTPHandler,
	}

	s.logger.Infof("Server started at: %s", addr)

	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		s.logger.WithError(err).Fatalf("HTTP server error")
	}
}
