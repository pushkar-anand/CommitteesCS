package main

import (
	"committees/config"
)

func main() {
	appConfig := config.GetAppConfig()
	config.SetupLogger()

	logger := config.GetLogger()

	server := NewServer(logger, appConfig)

	server.Initialize()

	//defer server.db.Conn.Close()

	server.Listen()

	<-server.connClose
	logger.Info("Shutdown complete")
}
