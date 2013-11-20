package go_tanks

import (
  "net"
  log "./log"
  "strconv"
)

type TCPServer struct {
  Port      int
  Address   string
  listener  *net.TCPListener
}

func (srv *TCPServer) run (channel chan<- *Client) {
  addrStr := srv.Address + ":"  + strconv.Itoa(srv.Port)
  log.Server("Try to listen tcp ", addrStr)

  addr, err := net.ResolveTCPAddr("tcp", addrStr)
  if ( err != nil ) { log.Fatal(err) } else { log.Server("TCP Server started on ", addr)}

  listener, err := net.ListenTCP("tcp", addr)
  srv.listener = listener
  if ( err != nil ) { log.Fatal(err) }


  for {
    conn, err := listener.Accept()
    if (err != nil ) { log.Fatal( err ) }
    channel <- NewClient(conn)
  }
}
