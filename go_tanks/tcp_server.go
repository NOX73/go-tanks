package go_tanks

import (
  "net"
  "log"
  "strconv"
)

type TCPServer struct {
  Port      int
  Address   string
  listener  *net.TCPListener
}

func (srv *TCPServer) run (c chan Client) {
  addrStr := srv.Address + ":"  + strconv.Itoa(srv.Port)
  log.Println("Try to listen tcp ", addrStr)

  addr, err := net.ResolveTCPAddr("tcp", addrStr)
  if ( err != nil ) { log.Panic(err) } else { log.Println("TCP Server started on ", addr)}

  listener, err := net.ListenTCP("tcp", addr)
  srv.listener = listener
  if ( err != nil ) { log.Panic(err) }

  
  for {
    conn, err := listener.Accept()
    if (err != nil ) { log.Panic( err ) } else { log.Pringln("New client conected") }
  }
}
