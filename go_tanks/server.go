package go_tanks

import(
  "log"
)

func NewServer(cfg Config) *Server {
  return &Server{config: &cfg}
}

type Server struct {
  config          *Config
  world           *World
  frontController *FrontController
};

func (srv *Server) Run () {
  srv.runWorld();

  log.Println("Server starting...");
  srv.run()
}

func (srv *Server) run () {
  tcpServer := TCPServer { Port: srv.config.Port, Address: srv.config.Address }
  channel := make(chan *Client)


  go tcpServer.run( channel )

  srv.frontController = &FrontController{ World: srv.world, ClientsChannel: channel }
  srv.frontController.Accept();

}

func (srv *Server) runWorld () {
  srv.world = &World{};
  srv.world.run();
}

