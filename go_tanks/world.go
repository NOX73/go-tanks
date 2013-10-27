package go_tanks

import (
  "time"
  "log"
)

type World struct {
  moment time.Time
}

func (w *World) run () {
  go w.start();
  log.Println("World started.");
}

func (w *World) start () {
  ticker := time.Tick(100 * time.Millisecond);
  for now := range ticker {
    w.moment = now
    //log.Println(now)
  }
}
