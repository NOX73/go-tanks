package go_tanks

import (
  "net"
  "github.com/gorilla/websocket"
  i "./interfaces"
  "encoding/json"
  "bufio"
)

const (
  EOL = "\n"
)

type IConn interface {
  ReadMessage()( *i.Message, error )
  WriteMessage( *i.Message ) error
  Close() error
  RemoteAddr() net.Addr
}

type NetConn struct {
  conn  *net.Conn
  r     *bufio.Reader
}

type WsConn struct {
  conn *websocket.Conn
}

// WebSocket Connection
func NewWsConn (conn *websocket.Conn) IConn {
  ws := WsConn{ conn: conn }
  return &ws
}

func ( c *WsConn ) ReadMessage () ( *i.Message, error ) {
  var mType int
  var buffer []byte
  var err error

  for mType != websocket.TextMessage {
    mType, buffer, err = c.conn.ReadMessage()
    if err != nil { return nil, err }
  }

  message := i.Message{}
  err = json.Unmarshal(buffer, &message)
  if( err != nil) { return nil, err }

  return &message, nil
}

func ( c *WsConn ) WriteMessage ( m *i.Message ) error {
  jsonStr, err := json.Marshal( m )

  if( err != nil ){ return err }

  return c.conn.WriteMessage(websocket.TextMessage, jsonStr)
}

func ( c *WsConn ) Close() error {
  return (*c.conn).Close()
}

func ( c *WsConn ) RemoteAddr() net.Addr {
  return (*c.conn).RemoteAddr()
}

// NET Connection
func NewNetConn (conn *net.Conn) IConn {
  nc := NetConn{ conn: conn, r: bufio.NewReader(*conn) }
  return &nc
}

func ( c *NetConn ) ReadMessage () ( *i.Message, error ) {
  buffer := []byte(nil)
  r := c.r

  for {
    part, prefix, err := r.ReadLine()
    if err != nil { return nil, err }

    if len(part) == 0 { continue }

    buffer = append( buffer, part... )

    for prefix && err == nil {
      part, prefix, err = r.ReadLine()
      if err != nil { return nil, err }
      buffer = append( buffer, part... )
    }

    break
  }

  message := i.Message{}
  err := json.Unmarshal(buffer, &message)
  if( err != nil) { return nil, err }

  return &message, nil

}

func ( c *NetConn )  WriteMessage ( m *i.Message ) error {
  jsonStr, err := json.Marshal( m )

  if( err != nil ){ return err }

  for len(jsonStr) > 0 {
    n, err := (*c.conn).Write(jsonStr)
    if err != nil { return err }

    jsonStr = jsonStr[n:]
  }

  (*c.conn).Write([]byte(EOL))

  return nil
}

func ( c *NetConn ) Close() error {
  return (*c.conn).Close()
}

func ( c *NetConn ) RemoteAddr() net.Addr {
  return (*c.conn).RemoteAddr()
}


