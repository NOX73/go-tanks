package go_tanks

import (
  "time"
)

type WorldConfig struct {
  TickDelay       time.Duration
  mapWidth        int
  mapHeight       int
  TankSpeed       float64
  TankRadius      int
  TankHealth      int
  GunSpeed        float64
  GunReload       int
  GunReloadSpeed  int
  GunMaxTemperature  int
  GunTemperatureWarming  int
  GunTemperatureCooling  int
  BulletSpeed     float64
}

type Config struct {
  Address         string
  Port            int
  World           *WorldConfig
}

var DefaultConfig = Config{
  Address:    "0.0.0.0",
  Port:       9292,
  World: &WorldConfig {
    TickDelay:  20,
    mapWidth:   1024,
    mapHeight:  768,

    //pixels per tick
    TankSpeed:  2,
    TankRadius: 10,
    TankHealth: 1,

    // grad per tick
    GunSpeed:  2,

    GunReload: 100,
    GunReloadSpeed: 5,

    GunMaxTemperature: 500,
    GunTemperatureWarming: 100,
    GunTemperatureCooling: 2,

    // puxels per tick
    BulletSpeed:  5,
  },
}
