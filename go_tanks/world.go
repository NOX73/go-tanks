package go_tanks

import (
  "time"
  log "./log"
  i "./interfaces"
)

const (
  NEW_TANK = iota
)

type Command struct {
  Type      int
  Channel   i.MessageChan
  Data      *i.Message
}

type World struct {
  Moment          time.Time
  TickDelay       time.Duration
  CommandChannel  chan *Command
}

func NewWorld (config *Config) *World {
  return &World{ TickDelay: config.TickDelay, CommandChannel: make(chan *Command, 5) };
}

func (w *World) run () {
  go w.start();
  log.Server("World started.");
}

func (w *World) start () {
  ticker := time.Tick( w.TickDelay * time.Millisecond );
  for now := range ticker {
    w.Moment = now
    w.processCommands()
  }
}

func ( w *World ) processCommands () {
  count := len( w.CommandChannel )
  for i := 0; i < count; i++ {
    command := <-w.CommandChannel
    w.processCommand( command )
  }
}

func ( w *World ) processCommand ( command *Command ) {
  switch command.Type {
  case NEW_TANK:
    w.addNewTank( command )
  }
}

func ( w *World ) NewTank ( channel i.MessageChan ) *i.Message {
  command := Command{ Type: NEW_TANK, Channel: channel }

  w.CommandChannel <- &command

  return <-channel
}

func ( w *World ) addNewTank ( command *Command ) {
  id := 1
  message := &i.Message{ "id": id }

  log.World("New Tank with id = ", id)
  command.Channel <- message
}
