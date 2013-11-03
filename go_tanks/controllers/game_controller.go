package go_tanks

import (
  //log "../log"
  //i "../interfaces"
)

type GameController struct {
  Controller
}

func ( c *GameController ) JoinToGame () error {
  c.addToWorld()

  c.World.NewTank( c.Client )
  message := *(c.Client.ReadInBox())

  message["message"] = "Your tank id"

  c.Client.SetTankId( message["id"].(int) )
  c.Client.SendMessage( &message )

  return nil
}

func ( c *GameController ) addToWorld () {
  c.World.AttachClient( c.Client )
}
