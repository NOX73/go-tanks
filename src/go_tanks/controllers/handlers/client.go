package go_tanks

import (
  //log "../../log"
  i "../../interfaces"
)

func ClientMessageHandler( w i.World, c i.Client, m *i.Message ) error {
  message := *m

  if message["WorldDisabled"] != nil {
    c.SetWorldDisabled(message["WorldDisabled"].(bool))
  }

  if message["WorldFrequency"] != nil {
    c.SetWorldFrequency(message["WorldFrequency"].(int))
  }

  return nil
}
