package go_tanks

type Message            map[string]interface{}
type MessageChan        chan *Message

type Client interface {
  SendMessage ( m *Message ) error
  ReadMessage () ( *Message, error )
  SetAuthCredentials ( login, password string)
  Channel () MessageChan
  SetTankId ( int )
  Login () *string
  Password () *string
}
