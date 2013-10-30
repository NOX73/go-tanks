package go_tanks

type Tank struct {
  Id      int
}

func NewTank ( id int ) *Tank {
  tank := Tank{ Id: id }
  return &tank
}
