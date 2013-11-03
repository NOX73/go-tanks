package go_tanks

type Tank struct {
  Id      int
  Coords  *Coords
}

func NewTank ( id int, coords *Coords ) *Tank {
  tank := Tank{ Id: id, Coords: coords }
  return &tank
}
