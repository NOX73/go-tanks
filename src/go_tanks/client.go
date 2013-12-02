package go_tanks

import (
  "github.com/gorilla/websocket"
  "net"
  i "./interfaces"
  "errors"
  log "./log"
)

const (
  NON_AUTHORIZED = iota
  AUTHORIZED
)

const (
  INBOX_CAPACITY = 5
  OUTBOX_CAPACITY = 5
  CLIENT_BUFFER_CAPACITY = 5
)

type Client struct {
  Connection  IConn
  State       int
  login       string
  password    string
  Tank        *Tank
  outBox      i.MessageChan
  inBox       i.MessageChan
  messageBox  i.MessageChan
  worldRecieveDisabled bool
}

func ( c *Client )  RemoteAddr () ( net.Addr ) {
  return c.Connection.RemoteAddr()
}

func NewClient ( conn IConn ) ( *Client ) {
  client := &Client{
    Connection: conn,
    State: NON_AUTHORIZED,
    inBox: make( i.MessageChan, INBOX_CAPACITY ),
    outBox: make( i.MessageChan, OUTBOX_CAPACITY ),
    messageBox: make( i.MessageChan, CLIENT_BUFFER_CAPACITY ),
    worldRecieveDisabled: false,
  }
  client.Init()
  return client
}

func NewWsClient ( ws *websocket.Conn ) ( *Client ) {
  return NewClient( NewWsConn(ws) )
}

func NewNetClient ( conn *net.Conn ) ( *Client ) {
  return NewClient( NewNetConn(conn) )
}

func ( c *Client ) Init () {
  go c.startSendMessageLoop()
}

func ( c *Client ) Close () {
  c.Connection.Close()
}

func ( c *Client ) SendMessage ( m *i.Message ) error {
  select {
  case c.messageBox <- m:
    return nil
  default:
    return errors.New("Slow client.")
  }
}

func ( c *Client ) startSendMessageLoop () {
  for message := range c.messageBox {
    c.Connection.WriteMessage( message )
  }
}

func ( c *Client ) ReadMessage () ( *i.Message, error ) {
  return c.Connection.ReadMessage()
}

func ( c *Client ) SetAuthCredentials ( login, password string ) {
  c.login = login
  c.password = password
}

func ( c *Client ) Login () *string {
  return &c.login
}

func ( c *Client ) Password () *string {
  return &c.password
}

func ( c *Client ) HasTank () bool {
  return c.Tank != nil
}

func ( c *Client ) SetTank ( tank interface{} ) {
  c.Tank = tank.(*Tank)
}

func ( c *Client ) GetTank () interface{} {
  return c.Tank
}

func ( c *Client ) InBox () i.MessageChan {
  return c.inBox
}

func ( c *Client ) OutBox () i.MessageChan {
  return c.outBox
}

func ( c *Client ) ReadInBox () *i.Message {
  return <-c.inBox
}

func ( c * Client ) WriteInBox ( m *i.Message ) {
  c.inBox <- m
}

func ( c *Client ) ReadOutBox () *i.Message {
  return <-c.outBox
}

func ( c * Client ) WriteOutBox ( m *i.Message )  {
  c.outBox <- m
}

func ( c *Client ) SendWorld ( m *i.Message ) {
  if( c.worldRecieveDisabled ) { return }

  err := c.SendMessage(m)
  if( err != nil ) { log.Error(err) }
}

func ( c *Client ) SetWorldRecieveDisabled ( val bool ) {
  c.worldRecieveDisabled = val
}

func ( c *Client ) OutBoxHasPlace () bool {
  return cap(c.outBox) - len(c.outBox) > 0
}
