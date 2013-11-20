package go_tanks

import (
  "code.google.com/p/go.net/websocket"
)

var ClientChannel = make(chan *Client)

func NewGoTanksWsClient ( ws *websocket.Conn ) {
  NewWsClient( ws )
}
