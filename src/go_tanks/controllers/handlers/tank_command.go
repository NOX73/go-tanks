package go_tanks

import (
  //log "../log"
  i "../../interfaces"
)

func TankCommandMessageHandler( w i.World, c i.Client, m *i.Message ) error {

  select {
  case c.OutBox() <- m:
  default:
    c.SendMessage( i.ErrorMessage( "You are sending messages too fast." ) )
  }

  return nil
}
