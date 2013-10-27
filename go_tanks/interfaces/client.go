package go_tanks

type Message map[string]interface{}
type Client interface {
  SendMessage ( m *Message ) error
  ReadMessage ( ) (*Message, error)
  SetAuthCredentials ( login, password string)
  Login () string
  Password () string
}
