package server

import (
	"net/http"
	"uangq-be/cmd/handler"
	"uangq-be/configs"
)

type Server struct {
	Router       *http.ServeMux
	ServerConfig configs.ServerConfig
}

var MyServer Server

func InitServer() {

	MyServer = Server{
		Router:       http.NewServeMux(),
		ServerConfig: configs.HttpServerConfig,
	}
	MyServer.InitRoutes()
}

func (s Server) InitRoutes() {
	MyServer.Router.HandleFunc("/login", handler.HandleLogin())
	MyServer.Router.HandleFunc("/signup", handler.HandleSignup())
}
