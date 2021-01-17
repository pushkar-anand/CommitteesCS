package main

import (
	"committees/config"
	"committees/template"
)

func main() {
	appConfig := config.GetAppConfig()
	config.SetupLogger()

	logger := config.GetLogger()

	server := NewServer(logger, appConfig)

	template.ParseTemplates()

	server.Initialize()

	//defer server.db.Conn.Close()

	server.Listen()

	<-server.connClose
	logger.Info("Shutdown complete")
}
