package go_tanks

type Message map[string]interface{}
type Client interface {
  SendMessage ( m *Message ) error
}
