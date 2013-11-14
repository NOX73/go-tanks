package go_tanks

import (
  "time"
)

type Config struct {
        Address         string
        Port            int
        TickDelay       time.Duration
        mapWidth        int
        mapHeight       int
        TickSpeed       int
}

var DefaultConfig = Config{
        Address:    "0.0.0.0",
        Port:       9292,
        TickDelay:  100,
        mapWidth:   1024,
        mapHeight:  768,
        TickSpeed:  20,
}
