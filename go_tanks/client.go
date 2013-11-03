package go_tanks

import (
  "net"
  "encoding/json"
  i "./interfaces"
  "bufio"
  "errors"
)

const (
  NON_AUTHORIZED = iota
  AUTHORIZED
  EOL = "\n"
)

const (
  INBOX_CAPACITY = 5
  OUTBOX_CAPACITY = 5
)

type Client struct {
  Connection  net.Conn
  State       int
  Reader      *bufio.Reader
  login       string
  password    string
  TankId      int
  outBox      i.MessageChan
  inBox       i.MessageChan
}

func (c *Client) RemoteAddr () (net.Addr) {
  return c.Connection.RemoteAddr()
}

func NewClient(conn net.Conn) (*Client) {
  return &Client{
    Connection: conn,
    State: NON_AUTHORIZED,
    Reader: bufio.NewReader(conn),
    inBox: make(i.MessageChan, INBOX_CAPACITY),
    outBox: make(i.MessageChan, OUTBOX_CAPACITY),
  }
}

func (c *Client) Close () {
  c.Connection.Close()
}

func (c *Client) SendMessage ( m *i.Message ) error {
  jsonStr, err := json.Marshal(m)
  if( err != nil ){ return err }

  c.Connection.Write(jsonStr)
  c.Connection.Write([]byte(EOL))

  return nil
}

func (c *Client) ReadMessage () ( *i.Message, error ) {
  buffer := []byte(nil)

  for {
    part, prefix, err := c.Reader.ReadLine()
    if err != nil { return nil, errors.New("Connection read error.") }

    if len(part) == 0 { continue }

    buffer = append(buffer, part...)

    for prefix && err == nil {
      part, prefix, err = c.Reader.ReadLine()
      if err != nil { return nil, errors.New("Connection read error.") }
      buffer = append(buffer, part...)
    }

    break
  }

  message := i.Message{}
  err := json.Unmarshal(buffer, &message)
  if( err != nil) { return nil, err }

  return &message, nil
}

func (c *Client) SetAuthCredentials ( login, password string) {
  c.login = login
  c.password = password
}

func (c *Client) Login () *string {
  return &c.login
}

func (c *Client) Password () *string {
  return &c.password
}

func (c *Client) SetTankId (id int) {
  c.TankId = id
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

