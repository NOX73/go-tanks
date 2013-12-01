package go_tanks

type Live struct {
  Tanks       map[int]*Tank
  Map         *Map
  TickSpeed   int
}

func NewLive ( config *Config ) *Live {
  return &Live {
    Tanks: make(map[int]*Tank),
    Map: NewMap( config ),
    TickSpeed: config.TickSpeed,
  }
}

func ( l *Live ) MoveTanksTick () {

  l.EachTanks ( func ( tank *Tank, _ int ) {
    coords, direction := tank.CalculateMove( l.TickSpeed )

    if( coords.X < 0 || coords.X > l.Map.Width ) { coords.X = tank.Coords.X }
    if( coords.Y < 0 || coords.Y > l.Map.Height ) { coords.Y = tank.Coords.Y }

    tank.ApplyMove( coords, direction )
  })

}

func ( l *Live ) EachTanks ( f func( *Tank, int ) ) {
  var i int
  for _, t := range l.Tanks { f(t, i); i++ }
}

func ( l *Live ) AddTank ( tank *Tank ) {
  l.Tanks[tank.Id] = tank
}
