package go_tanks

import (
  log "./log"
  controllers "./controllers"
  i "./interfaces"
  "fmt"
)

type Dispatcher struct {
  World   *World
  Client  *Client
}

func ( d *Dispatcher ) run () {
  d.dispatch()
}

func ( d *Dispatcher ) dispatch () error {
  var err error;

  err = d.authStep();
  if ( err != nil ) { return d.closeWithError( err ) }

  err = d.gameStep();
  if ( err != nil ) { return d.closeWithError( err ) }

  return nil
}

func ( d *Dispatcher ) gameStep () error {
  controller := controllers.GameController{ controllers.Controller { Client: d.Client, World: d.World } }
  return controller.JoinToGame()
}

func ( d *Dispatcher ) authStep () error {
  controller := controllers.AuthController{ controllers.Controller{ Client: d.Client, World: d.World } }
  return controller.Authorize()
}

func ( d *Dispatcher ) closeWithError (err error) error {
    log.Error(err)
    d.Client.SendMessage( &i.Message{ "message": fmt.Sprint( err ) } )
    d.Client.Close()

    return err
}
