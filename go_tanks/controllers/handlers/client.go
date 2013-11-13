package go_tanks

import (
  //log "../log"
  i "../../interfaces"
)

func ClientMessageHandler( w i.World, c i.Client, m *i.Message ) error {
  message := *m

  if message["WorldRecieveDisabled"] != nil {
    c.SetWorldRecieveDisabled(message["WorldRecieveDisabled"].(bool))
  }

  return nil
}
