package go_tanks

import (
  //log "../log"
  i "../../interfaces"
)

func TankCommandMessageHandler( w i.World, c i.Client, m *i.Message ) error {
  c.WriteOutBox( m )

  return nil
}
