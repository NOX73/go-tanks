package go_tanks

import (
  "math/rand"
)

type Map struct {
  Width   int
  Height  int
}

type Coords struct {
  X   float64
  Y   float64
}

func NewMap ( config *WorldConfig ) *Map {
  return &Map{
    Width: config.mapWidth,
    Height: config.mapHeight,
  }
}

func ( m* Map ) GetRandomCoords () *Coords {
  x := rand.Intn( m.Width )
  y := rand.Intn( m.Height )

  return &Coords{X: float64(x), Y: float64(y)}
}
