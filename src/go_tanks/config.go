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
        TankSpeed       float64
        TankRadius      int
        GunSpeed        float64
        BulletSpeed     float64
}

var DefaultConfig = Config{
        Address:    "0.0.0.0",
        Port:       9292,
        TickDelay:  100,
        mapWidth:   1024,
        mapHeight:  768,

        //pixels per tick
        TankSpeed:  10,
        TankRadius: 10,

        // grad per tick
        GunSpeed:  2,

         // puxels per tick
        BulletSpeed:  15,
}
