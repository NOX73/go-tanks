package go_tanks

import (
  //log "./log"
)

type Live struct {
  Tanks       map[int]*Tank
  Map         *Map
  TankSpeed   float64
  GunSpeed    float64
  BulletSpeed float64
  Bullets     []*Bullet
  ObjectIndex *ObjectIndex
}

func NewLive ( config *Config ) *Live {
  return &Live {
    Tanks: make(map[int]*Tank),
    Map: NewMap( config ),
    TankSpeed: config.TankSpeed,
    GunSpeed: config.GunSpeed,
    BulletSpeed: config.BulletSpeed,
    Bullets: make([]*Bullet, 0, 30),
    ObjectIndex: NewObjectIndex(),
  }
}

func ( l *Live ) MoveTanksTick () {

  l.EachTanks ( func ( tank *Tank, _ int ) {
    coords, direction := tank.CalculateMove( l.TankSpeed )

    l.ObjectIndex.ApplyTankPosition(coords, direction, tank, l.Map)

    tank.TurnGun( l.GunSpeed )
  })

}

func ( l *Live ) MoveBulletsTick () {
  l.EachBUllets ( func ( b *Bullet, _ int ) {
    coords, direction := b.CalculateMove( l.BulletSpeed )

    onMap := l.ObjectIndex.ValidateBulletPosition(coords, direction, b, l.Map)

    if !onMap {
      l.removeBullet( b )
    } else {
      b.ApplyMove( coords, direction )
    }

  })
}

func ( l *Live ) removeBullet ( bullet *Bullet ) {
  var i int
  var f bool
  var b *Bullet

  for i, b = range l.Bullets {
    if b == bullet { f = true; break }
  }

  if ( f ) { l.Bullets = append( l.Bullets[:i], l.Bullets[i+1:]... ) }
}

func ( l *Live ) RemoveTank ( tank *Tank ) {
  l.ObjectIndex.Remove( tank )
  delete( l.Tanks, tank.Id )
}

func ( l *Live ) remove ( tank *Tank ) {
}

func ( l *Live ) EachBUllets ( f func( *Bullet, int ) ) {
  var i int
  for _, b := range l.Bullets { f(b, i); i++ }
}

func ( l *Live ) EachTanks ( f func( *Tank, int ) ) {
  var i int
  for _, t := range l.Tanks { f(t, i); i++ }
}

func ( l *Live ) AddTank ( tank *Tank ) {
  l.ObjectIndex.Add( tank )
  l.Tanks[tank.Id] = tank
}
