package go_tanks

import (
  //log "../log"
  i "../interfaces"
  v "../validators"
  h "./handlers"
)

type GameController struct {
  Controller
}

func ( c *GameController ) JoinToGame () error {
  c.addToWorld()

  c.World.NewTank( c.Client )
  message := *( c.Client.ReadInBox() )

  message["Message"] = "Your tank id"

  c.Client.SetTankId( message["Id"].(int) )
  c.Client.SendMessage( &message )

  for { 
    message, err := c.Client.ReadMessage()
    if( err != nil ) { continue }

    err = v.ValidateUserMessage( message )

    if ( err != nil ) { c.Client.SendMessage( i.ErrorMessage( err.Error() ) ); continue }

    c.handleMessage( message )
  }

  return nil
}

func ( c *GameController ) addToWorld () {
  c.World.AttachClient( c.Client )
}

func ( c *GameController ) handleMessage ( m *i.Message ) error {
  return h.HandleMessage(c.World, c.Client, m)
}
