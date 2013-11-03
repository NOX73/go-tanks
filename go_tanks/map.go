package go_tanks

import (
  "math/rand"
)

type Map struct {
  Width   int
  Height  int
}

type Coords struct {
  X   int
  Y   int
}

func NewMap ( config *Config ) *Map {
  return &Map{
    Width: config.mapWidth,
    Height: config.mapHeight,
  }
}

func ( m* Map ) GetRandomCoords () *Coords {
  x := rand.Intn( m.Width )
  y := rand.Intn( m.Height )

  return &Coords{X: x, Y: y}
}
