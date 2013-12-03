package go_tanks

type Gun struct {
  Direction         float64
  Reload            int64
  Temperature       int64
  TurnAngle         float64
}

func ( g *Gun ) fire ( t *Tank ) *Bullet {
  return NewBullet( t );
}

