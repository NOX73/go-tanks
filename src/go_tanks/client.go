package go_tanks

import (
  "github.com/gorilla/websocket"
  "net"
  i "./interfaces"
  "errors"
  log "./log"
  "time"
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
  State         int
  login         string
  password      string
  Tank          *Tank
  // inside message channels
  outBox        i.MessageChan
  inBox         i.MessageChan
  // client message channels
  inClientBox   i.MessageChan
  outClientBox  i.MessageChan

  worldRecieveDisabled bool
  Closed        bool
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
    inClientBox: make( i.MessageChan, CLIENT_BUFFER_CAPACITY ),
    outClientBox: make( i.MessageChan, CLIENT_BUFFER_CAPACITY ),
    worldRecieveDisabled: false,
    Closed: false,
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
  go c.startReceiveMessageLoop()
}

func ( c *Client ) ReInit () {
  c.Tank = nil
  c.State = NON_AUTHORIZED
  c.login = ""
  c.password = ""
  log.Client("Client reinitialized.")
}

func ( c *Client ) Close () {
  c.Closed = true

  // Wait for write all messages
  <- time.After(time.Second)

  c.Connection.Close()
  close(c.outClientBox)

  log.Client("Client closed.")
}

func ( c *Client ) SendMessage ( m *i.Message ) error {
  select {
  case c.outClientBox <- m:
    return nil
  default:
    return errors.New("Slow client.")
  }
}

func ( c *Client ) startReceiveMessageLoop () {
  defer close(c.inClientBox)

  for {
    message, err := c.Connection.ReadMessage()
    if err != nil { break }

    select {
    case c.inClientBox <- message:
    default:
    }

  }

  log.Client("Receive message finished.")
}
func ( c *Client ) startSendMessageLoop () {

  for {
    message, ok := <- c.outClientBox
    if !ok { break }
    c.Connection.WriteMessage( message )
  }

  log.Client("Send message finished.")
}

func ( c *Client ) ReadMessage () ( *i.Message, error ) {
  return <- c.inClientBox, nil
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

func ( c *Client ) InClientBox () i.MessageChan {
  return c.inClientBox
}

func ( c *Client ) OutClientBox () i.MessageChan {
  return c.outClientBox
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
