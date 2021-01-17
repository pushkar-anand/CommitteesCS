package main

import (
	"committees/config"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var server *Server

func TestMain(m *testing.M) {
	appConfig := config.GetAppConfig()
	logger := config.GetLogger()

	server = NewServer(logger, appConfig)
	server.Initialize()

	code := m.Run()

	os.Exit(code)
}

func TestServer_Listen(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/", nil)

	rr := httptest.NewRecorder()
	server.router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
