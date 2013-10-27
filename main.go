package main

import (
  "./go_tanks"
)

func main() {
  server := go_tanks.NewServer(go_tanks.DefaultConfig)
  server.Run();
}

