package main

import (
  "./src/go_tanks"
  "math/rand"
  "time"
  "./src/web"
)

func main() {
  go runWeb( )
  runGoTanks(  )
}

func runWeb () {
  server := web.NewServer( go_tanks.NewGoTanksWsClient )
  server.Run()
}

func runGoTanks () {
  config := go_tanks.DefaultConfig

  rand.Seed( time.Now().UTC().UnixNano())
  server := go_tanks.NewServer( config )
  server.Run();
}
