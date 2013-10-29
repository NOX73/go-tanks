package go_tanks

import (
  "time"
)

type Config struct {
        Address         string
        Port            int
        TickDelay       time.Duration
}

var DefaultConfig = Config{
        Address:    "0.0.0.0",
        Port:       9292,
        TickDelay:  100,
}
