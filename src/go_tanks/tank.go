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

func ( t *Tank ) Move ( speed int ) {

  sumMotor := math.Min( t.LeftMotor , t.RightMotor )

  t.Coords.X += int( math.Cos( t.Direction ) * float64(speed) * sumMotor )
  t.Coords.Y += int( math.Sin( t.Direction ) * float64(speed) * sumMotor )

  rotationSpeed := t.LeftMotor - t.RightMotor

  t.Direction += rotationSpeed

  if t.Direction < 0 { t.Direction += 360 }
  if t.Direction > 360 { t.Direction -= 360 }

}

