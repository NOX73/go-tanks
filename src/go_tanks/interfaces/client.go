package go_tanks

const (
  NEW_TANK = iota
  NEW_CLIENT
  REMOVE_CLIENT
  TANK_COMMAND
)

type Message            map[string]interface{}
type MessageChan        chan *Message
type Tank interface {}

type Client interface {
  SendMessage ( *Message ) error
  ReadMessage () ( *Message, error )

  SetAuthCredentials ( login, password string)
  SetTank ( interface{} )
  GetTank () interface{}
  HasTank () bool

  Login () *string
  Password () *string

  OutBox () MessageChan
  InBox () MessageChan

  ReadInBox () *Message
  WriteInBox ( *Message )
  ReadOutBox () *Message
  WriteOutBox ( *Message )

  OutBoxHasPlace () bool

  SendWorld ( *Message )
  SetWorldRecieveDisabled( bool )
}

var typeToIdMaping = map[string]int {
  "TankCommand": TANK_COMMAND,
}

func ( m *Message ) GetTypeId () interface{} {
  message := *m
  if message["TypeId"] == nil {
    message["TypeId"] = typeToIdMaping[ message["Type"].(string) ]
  }

  return message["TypeId"].(int)
}

func ErrorMessage ( message string ) *Message {
  return &Message{"Type":"Error", "Message": message}
}
