package go_tanks

import (
  "errors"
  log "../log"
  interfaces "../interfaces"
  "fmt"
)

const (
  HELLO = "Hello! You should authorize befor join the game!"
)

type AuthController struct {
  Client  interfaces.Client
  World   interface {} 
}

func (c *AuthController) Authorize () error {
  c.sendHello();
  c.readAuth();
  return errors.New( fmt.Sprintf("Authorization failed with credentials: %s / %s", c.Client.Login(), c.Client.Password()) )
}

func ( c *AuthController ) sendHello () {
  c.Client.SendMessage(&interfaces.Message{
    "message":  HELLO,
  })
}

func ( c *AuthController ) readAuth () error {
  m, err := c.Client.ReadMessage()
  if err != nil { log.Error("Auth failed: ", err); return err }
  message := *m

  c.Client.SetAuthCredentials( message["login"].(string) , message["password"].(string) )

  return nil
}
