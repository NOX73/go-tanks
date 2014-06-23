package go_tanks

import (
	log "./log"
)

func NewServer(cfg Config) *Server {
	return &Server{config: &cfg}
}

type Server struct {
	config *Config
	world  *World
	rooms  *Rooms
}

func (srv *Server) Run() {

	log.Server("Server starting...")
	srv.run()
}

func (srv *Server) run() {
	tcpServer := TCPServer{Port: srv.config.Port, Address: srv.config.Address}
	channel := ClientChannel

	go tcpServer.run(channel)
	rooms := NewRooms(*srv.config.World)

	NewFrontController(channel, rooms).Accept()
}
