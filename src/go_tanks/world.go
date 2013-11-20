package go_tanks

import (
  "time"
  log "./log"
  i "./interfaces"
)

type World struct {
  TanksCounter    int
  Moment          time.Time
  TickDelay       time.Duration
  CommandChannel  i.MessageChan
  Clients         []i.Client
  Live            *Live
}

func NewWorld ( config *Config ) *World {

  world := &World{ 
    TickDelay: config.TickDelay,
    CommandChannel: make(i.MessageChan, 5),
    Live: NewLive( config ),
  };

  return world
}

func (w *World) run () {
  go w.start();
  log.Server("World started.");
}

func (w *World) start () {
  ticker := time.Tick( w.TickDelay * time.Millisecond );
  for now := range ticker {
    w.Moment = now
    w.processClientsCommands()
    w.processCommands()
    w.calculateWorld()
    w.sendWorldToClients()
    w.liveTick()
  }
}

func ( w *World ) Tanks () *map[int]*Tank {
  return &w.Live.Tanks
}

func ( w *World ) Map () *Map {
  return w.Live.Map
}

func ( w *World ) sendWorldToClients () {
  tanks := make([]*Tank, len(*w.Tanks()))

  w.Live.EachTanks ( func ( tank *Tank, n int ) {
    tanks[n] = tank
  })

  snapShot := &i.Message{
    "Type": "World",
    "Tanks": tanks,
    "Map": w.Map(),
  }

  for _, client := range w.Clients {
    client.SendWorld( snapShot );
  }
}
func ( w *World ) calculateWorld () {
}

func ( w *World ) processClientsCommands () {
  for _, client := range w.Clients  {
    select {
    case command := <-client.OutBox():
      w.processCommand( command, client )
    default:
    }
  }
}

func ( w *World ) processCommands () {
  count := len( w.CommandChannel )
  for i:=0; i < count; i++ {
    command := <- w.CommandChannel
    w.processCommand( command, nil )
  }
}

func ( w *World ) AttachClient ( client i.Client ) {
  w.CommandChannel <- &i.Message{"TypeId": i.NEW_CLIENT, "Client": client}
}

func ( w *World ) DetachClient ( client i.Client ) {
  w.CommandChannel <- &i.Message{"TypeId": i.REMOVE_CLIENT, "Client": client}
}

func ( w *World ) processCommand ( m *i.Message, client i.Client ) {
  message := *m 
  if( client == nil){ client = message["Client"].(i.Client) }

  switch message.GetTypeId() {
  case i.NEW_TANK:
    w.addNewTank( client )
  case i.REMOVE_CLIENT:
    w.removeClient( client )
  case i.NEW_CLIENT:
    w.addNewClient( client )
  case i.TANK_COMMAND:
    w.processTankCommand( m, client )
  }
}

func ( w *World ) processTankCommand ( m *i.Message, client i.Client ) {
  message := *m
  tank := client.GetTank().(*Tank)

  if message["LeftMotor"] != nil { tank.LeftMotor = message["LeftMotor"].(float64)  }
  if message["RightMotor"] != nil { tank.RightMotor = message["RightMotor"].(float64)  }
}

func ( w *World ) NewTank ( client i.Client ) {
  message := i.Message{"TypeId": i.NEW_TANK}
  client.WriteOutBox( &message )
}

func ( w *World ) nextTankId () int {
  w.TanksCounter++
  return w.TanksCounter
}

func ( w *World ) addNewTank ( client i.Client ) {
  id := w.nextTankId()
  coords := w.Map().GetRandomCoords()
  tank := NewTank(id, coords)

  w.Live.AddTank( tank )

  replay := i.Message{ "Tank": tank, "Type": "Tank" }

  log.World("New Tank with id = ", id)
  client.WriteInBox( &replay )
}

func ( w *World ) removeTank ( tank *Tank ) {
  delete( *w.Tanks(), tank.Id )
  log.World("Tank with id = ", tank.Id, " was removed.")
}

func ( w *World ) addNewClient ( client i.Client ) {
  w.Clients = append( w.Clients, client )
  log.World("Client added. Clients count = ", len( w.Clients ))
}


func ( w *World) removeClient ( client i.Client ) {
  index := -1;
  for i, c := range w.Clients {

    if c == client {
      index = i
      break
    }
  }

  if index < 0 { log.Warning("Client hasn't been removed."); return }

  if client.HasTank() {
    tank := client.GetTank().(*Tank)
    w.removeTank( tank ) 
  }

  w.Clients = append(w.Clients[:index], w.Clients[index + 1:]...)
  log.World("Client removed. Clients count = ", len( w.Clients ))
}

func ( w *World ) liveTick () {
  w.Live.MoveTanksTick()
}
