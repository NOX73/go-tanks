package go_tanks

import (
  "errors"
  interfaces "../interfaces"
)

type AuthController struct {
  Client  interfaces.Client
  World   interface {} 
}

func (c *AuthController) Authorize () error {
  c.sendHello();
  return errors.New("Authorization failed.")
}

func ( c *AuthController) sendHello () {
  c.Client.SendMessage(&interfaces.Message{
    "message":  "Hello",
  })
}
