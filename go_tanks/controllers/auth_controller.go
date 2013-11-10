package go_tanks

import (
  "errors"
  log "../log"
  i "../interfaces"
  "fmt"
)

const (
  HELLO = "Hello! You should authorize befor join the game!"
)

type AuthController struct {
  Controller
}

func ( c *AuthController ) Authorize () error {
  c.sendHello();
  c.readAuth();

  if ( !c.isAuthorized() ){
    return errors.New( fmt.Sprintf("Authorization failed with credentials: %s / %s", c.Client.Login(), c.Client.Password() ) )
  }

  return nil
}

func ( c *AuthController ) sendHello () {
  c.Client.SendMessage( &i.Message{
    "Message":  HELLO,
  })
}

func ( c *AuthController ) readAuth () error {
  m, err := c.Client.ReadMessage()
  if err != nil { log.Error("Auth failed: ", err); return err }
  message := *m

  login := message["Login"];
  password := message["Password"];

  if( login == nil ) { login = "" }
  if( password == nil ) { password = "" }

  c.Client.SetAuthCredentials( login.(string) , password.(string) )

  return nil
}

func ( c * AuthController ) isAuthorized () bool {
  return len( *c.Client.Login() ) > 0
}
