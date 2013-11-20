package main

import (
  "./src/go_tanks"
  "math/rand"
  "time"
  "github.com/robfig/revel"
  "github.com/robfig/revel/harness"
)

func main() {
  go runWeb( )
  runGoTanks(  )
}

func runWeb () {
  revel.Init("dev", "web", "")
  revel.LoadMimeConfig()

  app, _ := harness.Build()
  app.Cmd().Run()
}

func runGoTanks () {
  config := go_tanks.DefaultConfig

  rand.Seed( time.Now().UTC().UnixNano())
  server := go_tanks.NewServer( config )
  server.Run();
}

