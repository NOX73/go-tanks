package go_tanks

import (
  "log"
)

type FrontController struct {
  World           *World
  ClientsChannel  <-chan *Client
}

func (fc *FrontController) Accept () {
  for {
    client := <- fc.ClientsChannel
    log.Println("CLIENT:\t New client connected (", client.RemoteAddr(), ")");

    fc.processClient(client)
  }
}

func (fc *FrontController) processClient (client *Client) {
  dispatcher := Dispatcher{Client: client, World: fc.World}
  go dispatcher.run()
}
