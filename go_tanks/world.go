package go_tanks

import (
  "time"
  log "./log"
  i "./interfaces"
)

const (
  NEW_TANK = iota
  NEW_CLIENT
)

type World struct {
  Map             *Map
  TanksCounter    int
  Moment          time.Time
  TickDelay       time.Duration
  CommandChannel  i.MessageChan
  Tanks           map[int]*Tank
  Clients         []i.Client
}

func NewWorld (config *Config) *World {
  return &World{ 
    TickDelay: config.TickDelay,
    CommandChannel: make(i.MessageChan, 5),
    Tanks: make( map[int]*Tank ),
    Map: NewMap(config),
  };
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
  }
}

func ( w *World ) sendWorldToClients () {
  tanks := make([]*Tank, len(w.Tanks))

  n := 0
  for _, tank := range w.Tanks {
    tanks[n] = tank
    n++
  }

  snapShot := &i.Message{
    "Tanks": tanks,
    "Map": w.Map,
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
  w.CommandChannel <- &i.Message{"Type": NEW_CLIENT, "Client": client}
}

func ( w *World ) processCommand ( command *i.Message, client i.Client ) {
  switch command.GetType() {
  case NEW_TANK:
    w.addNewTank( client )
  case NEW_CLIENT:
    w.addNewClient( (*command)["Client"].(i.Client) )
  }
}

func ( w *World ) NewTank ( client i.Client ) {
  message := i.Message{"Type": NEW_TANK}
  client.WriteOutBox( &message )
}

func ( w *World ) nextTankId () int {
  w.TanksCounter++
  return w.TanksCounter
}

func ( w *World ) addNewTank ( client i.Client ) {
  id := w.nextTankId()
  coords := w.Map.GetRandomCoords()
  tank := NewTank(id, coords)
  w.Tanks[id] = tank

  replay := i.Message{ "Id": id, "Tank": tank }

  log.World("New Tank with id = ", id)
  client.WriteInBox( &replay )
}

func ( w *World ) addNewClient ( client i.Client ) {
  log.World("New client.")
  w.Clients = append( w.Clients, client )
}
