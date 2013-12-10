package go_tanks

import (
  "math"
  //log "./log"
)

type Tank struct {
  Id            int
  Coords        *Coords
  Direction     float64
  LeftMotor     float64
  RightMotor    float64
  Gun           *Gun
}

func NewTank ( id int, coords *Coords ) *Tank {
  tank := Tank{
    Id: id,
    Coords: coords,
    LeftMotor: 0,
    RightMotor: 0,
    Direction: 0,
    Gun: &Gun{ Direction: 0, TurnAngle: 0 },
  }
  return &tank
}

func ( t *Tank ) CalculateMove ( speed float64 ) (*Coords, float64) {

  sumMotor := math.Min( t.LeftMotor , t.RightMotor )

  radDirection := (math.Pi * t.Direction) / 180
  x := t.Coords.X + int( math.Cos( radDirection ) * speed * sumMotor )
  y := t.Coords.Y + int( math.Sin( radDirection ) * speed * sumMotor )

  rotationSpeed := t.LeftMotor - t.RightMotor

  direction := t.Direction + rotationSpeed * speed

  if direction < 0 { direction += 360 }
  if direction > 360 { direction -= 360 }

  return &Coords{X: x, Y: y}, direction
}

func ( t *Tank ) ApplyMove ( c *Coords, d float64 ) {
  t.Coords = c
  t.Direction = d
}

func ( t *Tank ) Fire () *Bullet {
  return t.Gun.fire( t );
}

func ( t *Tank ) TurnGun ( speed float64 ) {
  if ( t.Gun.TurnAngle < 0.1 && t.Gun.TurnAngle > -0.1 ){ return }

  diff := t.Gun.TurnAngle

  if diff > 0 {
    diff = math.Min(speed, diff)
  } else {
    diff = math.Max(-speed, diff)
  }

  t.Gun.TurnAngle -= diff
  t.Gun.Direction += diff 

  if t.Gun.Direction < 0 { t.Gun.Direction += 360 }
  if t.Gun.Direction > 360 { t.Gun.Direction -= 360 }
}

func ( t *Tank ) GetCoords () *Coords {
  return t.Coords
}
