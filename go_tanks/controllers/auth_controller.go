package go_tanks

import (
  "errors"
  log "../log"
  i "../interfaces"
  "fmt"
  v "../validators"
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
    "Type": "Auth",
  })
}

func ( c *AuthController ) readAuth () error {
  m, err := c.Client.ReadMessage()
  if err != nil { log.Error("Auth failed: ", err); return err }

  err = v.ValidateAuthForm(m)
  if( err != nil ) { c.Client.SendMessage( &i.Message{ "Type": "Error", "Message": err.Error() } ); return err }

  message := *m
  c.Client.SetAuthCredentials( message["Login"].(string) , message["Password"].(string) )

  return nil
}

func ( c * AuthController ) isAuthorized () bool {
  return len( *c.Client.Login() ) > 0
}
