package go_tanks

import (
  //"errors"
  "math"
)

type Gun struct {
  Direction         float64
  Config            *WorldConfig `json:"-"`

  ReloadProgress    int
  Temperature       int
  TurnAngle         float64 `json:"-"`
}

func NewGun (config *WorldConfig) *Gun {
  return &Gun{ 
    Config: config,
  }
}

func ( g *Gun ) Fire ( tank *Tank ) *Bullet {
  if(g.ReloadProgress != 0 || g.Temperature > g.Config.GunMaxTemperature){
    return nil
  } else {
    g.ReloadProgress = g.Config.GunReload
    g.Temperature += g.Config.GunTemperatureWarming
    return NewBullet( tank, g.Config  )
  }
}

func ( g *Gun ) TickParams () {

  if g.ReloadProgress > 0 {
    g.ReloadProgress -= g.Config.GunReloadSpeed
    if g.ReloadProgress < 0 { g.ReloadProgress = 0 } 
  }

  if g.Temperature > 0 {
    g.Temperature -= g.Config.GunTemperatureCooling
    if g.Temperature < 0 { g.Temperature = 0 } 
  }

  g.Turn()
}

func ( g *Gun ) Turn () {
  speed := g.Config.GunSpeed

  if ( g.TurnAngle < 0.1 && g.TurnAngle > -0.1 ){ return }

  diff := g.TurnAngle

  if diff > 0 {
    diff = math.Min(speed, diff)
  } else {
    diff = math.Max(-speed, diff)
  }

  g.TurnAngle -= diff
  g.Direction += diff

  if g.Direction < 0 { g.Direction += 360 }
  if g.Direction > 360 { g.Direction -= 360 }

}
