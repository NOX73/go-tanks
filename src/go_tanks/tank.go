package go_tanks

import (
  "math"
)

type Tank struct {
  Id            int
  Coords        *Coords
  Direction     float64
  LeftMotor     float64
  RightMotor    float64
}

func NewTank ( id int, coords *Coords ) *Tank {
  tank := Tank{
    Id: id,
    Coords: coords,
    LeftMotor: 0,
    RightMotor: 0,
    Direction: 0,
  }
  return &tank
}

func ( t *Tank ) CalculateMove ( speed int ) (*Coords, float64) {

  sumMotor := math.Min( t.LeftMotor , t.RightMotor )

  radDirection := (math.Pi * t.Direction) / 180
  x := t.Coords.X + int( math.Cos( radDirection ) * float64(speed) * sumMotor )
  y := t.Coords.Y + int( math.Sin( radDirection ) * float64(speed) * sumMotor )

  rotationSpeed := t.LeftMotor - t.RightMotor

  direction := t.Direction + rotationSpeed * float64(speed)

  if direction < 0 { direction += 360 }
  if direction > 360 { direction -= 360 }

  return &Coords{X: x, Y: y}, direction 
}

func ( t *Tank ) ApplyMove ( c *Coords, d float64 ) {
  t.Coords = c
  t.Direction = d
}
