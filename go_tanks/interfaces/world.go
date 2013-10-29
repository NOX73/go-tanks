package go_tanks

type World interface {
  NewTank ( channel MessageChan ) *Message
}
