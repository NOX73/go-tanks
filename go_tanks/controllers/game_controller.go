package go_tanks

import (
  //log "../log"
  //i "../interfaces"
)

type GameController struct {
  Controller
}

func ( c *GameController ) JoinToGame () error {

  ch := c.Client.Channel()

  m := c.World.NewTank( ch )
  message := *m

  //message := i.Message{"id":3}

  message["message"] = "Your tank id"

  c.Client.SetTankId( message["id"].(int) )
  c.Client.SendMessage( &message )

  return nil
}
