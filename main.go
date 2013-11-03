package main

import (
  "./go_tanks"
  "math/rand"
  "time"
)

func main() {
  rand.Seed( time.Now().UTC().UnixNano())
  server := go_tanks.NewServer(go_tanks.DefaultConfig)
  server.Run();
}

