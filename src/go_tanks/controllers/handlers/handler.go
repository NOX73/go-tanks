package go_tanks

import (
  //log "../log"
  i "../../interfaces"
)

var messageHandlers = map[string]func( i.World, i.Client, *i.Message )error{
  "Client": ClientMessageHandler,
  "TankCommand": TankCommandMessageHandler,
  "Ping": PingMessageHandler,
}

func HandleMessage( w i.World, c i.Client, m *i.Message ) error {
  message := *m

  return messageHandlers[message["Type"].(string)](w, c, m)
}
