package go_tanks

type World interface {
  NewTank ( client Client )
  AttachClient ( client Client )
  DetachClient ( client Client )
}
