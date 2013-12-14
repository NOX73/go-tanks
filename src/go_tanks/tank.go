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
  Radius        int
  Health        int
  Config        *WorldConfig `json:"-"`
}

func NewTank ( id int, coords *Coords, config *WorldConfig ) *Tank {
  tank := Tank{
    Id: id,
    Coords: coords,
    LeftMotor: 0,
    RightMotor: 0,
    Direction: 0,
    Gun: NewGun( config ),
    Radius: config.TankRadius,
    Health: config.TankHealth,
    Config: config,
  }
  return &tank
}

func ( t *Tank ) CalculateMove () (*Coords, float64) {
  speed := t.Config.TankSpeed

  sumMotor := math.Min( t.LeftMotor , t.RightMotor )

  var x,y float64

  if t.LeftMotor * t.RightMotor > 0 {
    radDirection := (math.Pi * t.Direction) / 180
    x = t.Coords.X + math.Cos( radDirection ) * speed * sumMotor 
    y = t.Coords.Y + math.Sin( radDirection ) * speed * sumMotor 
  } else {
    x = t.Coords.X
    y = t.Coords.Y
  }

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

func ( t *Tank ) TickParams () {
  t.Gun.TickParams()
}

func ( t *Tank ) Fire () *Bullet {
  return t.Gun.Fire( t );
}

func ( t *Tank ) GetCoords () *Coords {
  return t.Coords
}

func ( t *Tank ) GetId () int {
  return t.Id
}

func ( t *Tank ) GetRadius () int {
  return t.Radius
}
