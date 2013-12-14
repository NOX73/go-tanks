package go_tanks

import (
  //log "./log"
)

type Live struct {
  Config      *WorldConfig
  Tanks       map[int]*Tank
  Map         *Map
  Bullets     []*Bullet
  ObjectIndex *ObjectIndex
}

func NewLive ( config *WorldConfig ) *Live {
  return &Live {
    Tanks: make(map[int]*Tank),
    Map: NewMap( config ),
    Config: config,
    Bullets: make([]*Bullet, 0, 30),
    ObjectIndex: NewObjectIndex(float64(config.TankRadius * 2)),
  }
}

func ( l *Live ) MoveTanksTick () {

  l.EachTanks ( func ( tank *Tank, _ int ) {
    coords, direction := tank.CalculateMove()

    l.ObjectIndex.ApplyTankPosition(coords, direction, tank, l.Map)
    tank.TickParams()
  })

}

func ( l *Live ) MoveBulletsTick () (hits []*Bullet) {
  var forRemove = make([]*Bullet,0,5)
  hits = make([]*Bullet,0,5)

  l.EachBUllets ( func ( b *Bullet, _ int ) {
    coords, direction := b.CalculateMove()

    tankHit, onMap := l.ObjectIndex.ValidateBulletPosition(coords, direction, b, l.Map)

    if !onMap {
      forRemove = append(forRemove, b)
    } else if tankHit == nil{
      b.ApplyMove( coords, direction )
    } else {
      b.HitTo(tankHit)
      hits = append(hits, b)
    }

  })

  for _, b := range forRemove {
    l.removeBullet( b )
  }

  return hits
}

func ( l *Live ) removeBullet ( bullet *Bullet ) {
  var i int
  var f bool
  var b *Bullet

  for i, b = range l.Bullets {
    if b == bullet { f = true; break }
  }

  l.ObjectIndex.Remove( bullet )
  if ( f ) { l.Bullets = append( l.Bullets[:i], l.Bullets[i+1:]... ) }
}

func ( l *Live ) RemoveTank ( tank *Tank ) {
  l.ObjectIndex.Remove( tank )
  delete( l.Tanks, tank.Id )
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

func ( l *Live ) AddBullet ( bullet *Bullet ) {
  l.ObjectIndex.Add( bullet )
  l.Bullets = append(l.Bullets, bullet)
}
