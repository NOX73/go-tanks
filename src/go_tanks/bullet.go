package go_tanks

import (
  "math"
)

type Bullet struct {
  Id            int
  TankId        int
  Tank          *Tank `json:"-"`
  Coords        *Coords
  Direction     float64
  HitToTank     *Tank `json:"-"`
  Speed         float64
}

func NewBullet ( tank *Tank, config *WorldConfig ) *Bullet {
  direction := tank.Direction + tank.Gun.Direction

  if direction < 0 { direction += 360 }
  if direction > 360 { direction -= 360 }

  coords := &Coords{ tank.Coords.X, tank.Coords.Y }

  radDirection := (math.Pi * direction) / 180
  coords.X += math.Cos( radDirection ) * float64(tank.Radius)
  coords.Y += math.Sin( radDirection ) * float64(tank.Radius)

  return &Bullet{
    TankId: tank.Id,
    Tank: tank,
    Coords: coords,
    Direction: direction,
    Speed: config.BulletSpeed,
  }
}

func ( b *Bullet ) CalculateMove () (*Coords, float64) {

  radDirection := (math.Pi * b.Direction) / 180
  x := b.Coords.X + math.Cos( radDirection ) * b.Speed
  y := b.Coords.Y + math.Sin( radDirection ) * b.Speed

  return &Coords{X: x, Y: y}, b.Direction
}

func ( b *Bullet ) ApplyMove ( c *Coords, d float64 ) {
  b.Coords = c
  b.Direction = d
}

func ( b *Bullet ) GetCoords () *Coords {
  return b.Coords
}

func ( b *Bullet ) GetRadius () int {
  return 1
}

func ( b *Bullet ) HitTo (tank *Tank) {
  b.HitToTank = tank
}

