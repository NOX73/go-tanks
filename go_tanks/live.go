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
    tank.Move( l.TickSpeed )
  })

}

func ( l *Live ) EachTanks ( f func( *Tank, int ) ) {
  var i int
  for _, t := range l.Tanks { f(t, i); i++ }
}

func ( l *Live ) AddTank ( tank *Tank ) {
  l.Tanks[tank.Id] = tank
}
