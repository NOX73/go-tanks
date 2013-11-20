package main

import (
  "./src/go_tanks"
  "math/rand"
  "time"
	//"reflect"
  //"github.com/robfig/revel"
  //"github.com/robfig/revel/harness"
  "./src/web"
  //"code.google.com/p/go.net/websocket"
)

func main() {
  go runWeb( )
  runGoTanks(  )
}

func runWeb () {
  server := web.NewServer()
  server.Run()
}

func runGoTanks () {
  config := go_tanks.DefaultConfig

  rand.Seed( time.Now().UTC().UnixNano())
  server := go_tanks.NewServer( config )
  server.Run();
}
