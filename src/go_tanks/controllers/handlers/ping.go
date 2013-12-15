package go_tanks

import (
  i "../../interfaces"
)


func PingMessageHandler( w i.World, c i.Client, m *i.Message ) error {
  message := *m

  c.SendMessage( &i.Message{ "Type":"Pong", "PongId": message["PingId"] } )

  return nil
}


