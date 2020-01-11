package main

import (
	"net/http"
	"uangq-be/bridge/db"
	"uangq-be/cmd/server"
	"uangq-be/common/logger"
)

func main() {
	// Init DB
	db.InitDB()

	// Init Server
	server.InitServer()

	logger.Log("Listening on " + server.MyServer.ServerConfig.Address)
	err := http.ListenAndServe(server.MyServer.ServerConfig.Address, server.MyServer.Router)
	if err != nil {
		logger.Log(err.Error())
	}
}
