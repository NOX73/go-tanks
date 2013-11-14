package go_tanks

type Tank struct {
  Id            int
  Coords        *Coords
  LeftMotor     float64
  RightMotor    float64
}

func NewTank ( id int, coords *Coords ) *Tank {
  tank := Tank{
    Id: id,
    Coords: coords,
    LeftMotor: 0,
    RightMotor: 0,
  }
  return &tank
}
