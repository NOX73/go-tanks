package go_tanks

import (
  "net"
  "encoding/json"
  i "./interfaces"
  "bufio"
  "errors"
  log "./log"
)

const (
  NON_AUTHORIZED = iota
  AUTHORIZED
  EOL = "\n"
)

const (
  INBOX_CAPACITY = 5
  OUTBOX_CAPACITY = 5
  CLIENT_BUFFER_CAPACITY = 5
)

type Client struct {
  Connection  net.Conn
  State       int
  Reader      *bufio.Reader
  login       string
  password    string
  Tank        *Tank
  outBox      i.MessageChan
  inBox       i.MessageChan
  jsonBox     chan *[]byte
  worldRecieveDisabled bool
}

func ( c *Client )  RemoteAddr () ( net.Addr ) {
  return c.Connection.RemoteAddr()
}

func NewClient ( conn net.Conn ) ( *Client ) {
  client := &Client{
    Connection: conn,
    State: NON_AUTHORIZED,
    Reader: bufio.NewReader(conn),
    inBox: make( i.MessageChan, INBOX_CAPACITY ),
    outBox: make( i.MessageChan, OUTBOX_CAPACITY ),
    jsonBox: make( chan *[]byte, CLIENT_BUFFER_CAPACITY ),
    worldRecieveDisabled: false,
  }
  client.Init()
  return client
}

func ( c *Client ) Init () {
  go c.startSendJsonLoop()
}

func ( c *Client ) Close () {
  c.Connection.Close()
}

func ( c *Client ) SendMessage ( m *i.Message ) error {
  jsonStr, err := json.Marshal( m )
  if( err != nil ){ return err }

  err = c.sendJson( &jsonStr )
  if(err != nil) {
    log.Client(err)
    return err 
  }

  return nil
}

func ( c *Client ) sendJson ( json *[]byte ) error {

  select {
    case c.jsonBox <- json:
      return nil
    default:
      return errors.New("Slow client.")
  }

  return nil
}

func ( c *Client ) startSendJsonLoop () {
  for jsonStr := range c.jsonBox {
    c.Connection.Write( *jsonStr )
    c.Connection.Write( []byte(EOL) )
  }
}

func ( c *Client ) ReadMessage () ( *i.Message, error ) {
  buffer := []byte(nil)

  for {
    part, prefix, err := c.Reader.ReadLine()
    if err != nil { return nil, err }

    if len(part) == 0 { continue }

    buffer = append( buffer, part... )

    for prefix && err == nil {
      part, prefix, err = c.Reader.ReadLine()
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
