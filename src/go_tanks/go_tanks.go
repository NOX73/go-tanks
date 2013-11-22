package go_tanks

import (
  "github.com/gorilla/websocket"
)

var ClientChannel = make(chan *Client)

func NewGoTanksWsClient ( ws *websocket.Conn ) {
  ClientChannel <- NewWsClient( ws )
}
