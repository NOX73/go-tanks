package go_tanks

import (
  "net"
  "encoding/json"
  interfaces "./interfaces"
)

const (
  NON_AUTHORIZED = iota
  AUTHORIZED
  EOL = "\n"
)

type Client struct {
  Connection  net.Conn
  State       int
}

func (c *Client) RemoteAddr () (net.Addr) {
  return c.Connection.RemoteAddr()
}

func NewClient(conn net.Conn) (*Client) {
  return &Client{Connection: conn, State: NON_AUTHORIZED}
}

func (c *Client) Close () {
  c.Connection.Close()
}

func (c *Client) SendMessage ( m *interfaces.Message ) error {
  jsonStr, err := json.Marshal(m)
  c.Connection.Write(jsonStr)
  c.Connection.Write([]byte(EOL))

  return err
}
